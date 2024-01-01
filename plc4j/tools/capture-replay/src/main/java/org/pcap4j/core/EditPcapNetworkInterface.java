/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package org.pcap4j.core;

import com.sun.jna.*;
import com.sun.jna.ptr.PointerByReference;
import org.pcap4j.util.ByteArrays;
import org.pcap4j.util.LinkLayerAddress;
import org.pcap4j.util.MacAddress;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public final class EditPcapNetworkInterface {

    private static final Logger logger = LoggerFactory.getLogger(PcapNetworkInterface.class);

    private static final int PCAP_IF_LOOPBACK = 0x00000001;
    private static final int PCAP_IF_UP = 0x00000002;
    private static final int PCAP_IF_RUNNING = 0x00000004;

    private final String name;
    private final String description;
    private final List<PcapAddress> addresses = new ArrayList<PcapAddress>();
    private final List<LinkLayerAddress> linkLayerAddresses = new ArrayList<LinkLayerAddress>();
    private final boolean loopBack;
    private final boolean up;
    private final boolean running;
    private final boolean local;

    private EditPcapNetworkInterface(NativeMappings.pcap_if pif, boolean local) {
        this.name = pif.name;
        this.description = pif.description;

        for (NativeMappings.pcap_addr pcapAddr = pif.addresses; pcapAddr != null; pcapAddr = pcapAddr.next) {
            if (pcapAddr.addr == null
                    && pcapAddr.netmask == null
                    && pcapAddr.broadaddr == null
                    && pcapAddr.dstaddr == null) {
                logger.warn("Empty pcap_addr on {} ({}). Ignore it.", name, description);
                continue;
            }

            short sa_family =
                    pcapAddr.addr != null
                            ? pcapAddr.addr.getSaFamily()
                            : pcapAddr.netmask != null
                                    ? pcapAddr.netmask.getSaFamily()
                                    : pcapAddr.broadaddr != null
                                            ? pcapAddr.broadaddr.getSaFamily()
                                            : pcapAddr.dstaddr != null
                                                    ? pcapAddr.dstaddr.getSaFamily()
                                                    /* default */ : Inets.AF_UNSPEC; // Never get here.
            if (sa_family == Inets.AF_INET) {
                addresses.add(PcapIpV4Address.newInstance(pcapAddr, sa_family, name));
            } else if (sa_family == Inets.AF_INET6) {
                addresses.add(PcapIpV6Address.newInstance(pcapAddr, sa_family, name));
            } else {
                if (Platform.isLinux() && sa_family == Inets.AF_PACKET) {
                    NativeMappings.sockaddr_ll sll = new NativeMappings.sockaddr_ll(pcapAddr.addr.getPointer());
                    byte[] addr = sll.sll_addr;
                    int addrLength = sll.sll_halen & 0xFF;
                    if (addrLength == 6) {
                        linkLayerAddresses.add(ByteArrays.getMacAddress(addr, 0));
                    } else if (addr.length == 0) {
                        continue;
                    } else {
                        // addrLength (i.e. sll_halen) may exceed addr.length (i.e. sll.sll_addr.length).
                        // If it's the case, need to shorten addrLength here to avoid
                        // ArrayIndexOutOfBoundsException in
                        // the succeeding ByteArrays.getSubArray().
                        // https://github.com/kaitoy/pcap4j/issues/228
                        addrLength = addrLength <= addr.length ? addrLength : addr.length;
                        linkLayerAddresses.add(
                                LinkLayerAddress.getByAddress(ByteArrays.getSubArray(addr, 0, addrLength)));
                    }
                } else if ((Platform.isMac() || Platform.isFreeBSD() || Platform.isOpenBSD())
                        || Platform.iskFreeBSD() && sa_family == Inets.AF_LINK) {
                    NativeMappings.sockaddr_dl sdl = new NativeMappings.sockaddr_dl(pcapAddr.addr.getPointer());
                    byte[] addr = sdl.getAddress();
                    if (addr.length == 6) {
                        linkLayerAddresses.add(MacAddress.getByAddress(addr));
                    } else if (addr.length == 0) {
                        continue;
                    } else {
                        linkLayerAddresses.add(LinkLayerAddress.getByAddress(addr));
                    }
                } else {
                    logger.warn("{} is not supported address family. Ignore it.", sa_family);
                }
            }
        }

        if (Platform.isWindows()) {
            MacAddress mac = getMacAddress(name);
            if (mac != null) {
                linkLayerAddresses.add(mac);
            }
        }

        this.loopBack = (pif.flags & PCAP_IF_LOOPBACK) != 0;
        this.up = (pif.flags & PCAP_IF_UP) != 0;
        this.running = (pif.flags & PCAP_IF_RUNNING) != 0;
        this.local = local;
    }

    static EditPcapNetworkInterface newInstance(NativeMappings.pcap_if pif, boolean local) {
        return new EditPcapNetworkInterface(pif, local);
    }

    /** @return name */
    public String getName() {
        return name;
    }

    /** @return description */
    public String getDescription() {
        return description;
    }

    /** @return inet addresses */
    public List<PcapAddress> getAddresses() {
        return new ArrayList<PcapAddress>(addresses);
    }

    /** @return link layer addresses */
    public ArrayList<LinkLayerAddress> getLinkLayerAddresses() {
        return new ArrayList<LinkLayerAddress>(linkLayerAddresses);
    }

    /**
     * Returns if this network interface is loopback. This method may always return false on some
     * environments.
     *
     * @return true if the network interface represented by this object is a loop back interface;
     *     false otherwise.
     */
    public boolean isLoopBack() {
        return loopBack;
    }

    /**
     * Returns if this network interface is up. This method may always return false on some
     * environments.
     *
     * @return true if the network interface represented by this object is up; false otherwise.
     */
    public boolean isUp() {
        return up;
    }

    /**
     * Returns if this network interface is running. This method may always return false on some
     * environments.
     *
     * @return true if the network interface represented by this object is running; false otherwise.
     */
    public boolean isRunning() {
        return running;
    }

    /**
     * @return true if the network interface represented by this object is a local interface; false
     *     otherwise.
     */
    public boolean isLocal() {
        return local;
    }

    /**
     * @param snaplen Snapshot length, which is the number of bytes captured for each packet.
     * @param mode mode
     * @param timeoutMillis Read timeout. Most OSs buffer packets. The OSs pass the packets to Pcap4j
     *     after the buffer gets full or the read timeout expires. Must be non-negative. May be
     *     ignored by some OSs. 0 means disable buffering on Solaris. 0 means infinite on the other
     *     OSs. 1 through 9 means infinite on Solaris.
     * @return a new PcapHandle object.
     * @throws PcapNativeException if an error occurs in the pcap native library.
     */
    public EditPcapHandle openLive(int snaplen, PcapNetworkInterface.PromiscuousMode mode, int timeoutMillis)
            throws PcapNativeException {
        if (mode == null) {
            StringBuilder sb = new StringBuilder();
            sb.append("mode: ").append(mode);
            throw new NullPointerException(sb.toString());
        }

        NativeMappings.PcapErrbuf errbuf = new NativeMappings.PcapErrbuf();
        Pointer handle =
                NativeMappings.pcap_open_live(name, snaplen, mode.getValue(), timeoutMillis, errbuf);
        if (handle == null || errbuf.length() != 0) {
            throw new PcapNativeException(errbuf.toString());
        }

        if (timeoutMillis == 0 && Platform.isSolaris()) {
            // disable buffering
            NativeMappings.timeval to = new NativeMappings.timeval();
            to.tv_sec = new NativeLong(0);
            to.tv_usec = new NativeLong(0);

            int rc =
                    NativeMappings.PcapLibrary.INSTANCE.strioctl(
                            NativeMappings.getFdFromPcapT(handle),
                            NativeMappings.SBIOCSTIME,
                            to.size(),
                            to.getPointer());

            if (rc < 0) {
                throw new PcapNativeException(
                        "SBIOCSTIME: "
                                + NativeMappings.pcap_strerror(NativeMappings.ERRNO_P.getInt(0)).getString(0));
            }
        }

        return new EditPcapHandle(handle, PcapHandle.TimestampPrecision.MICRO);
    }

    private MacAddress getMacAddress(String nifName) {
        Pointer lpAdapter = NativePacketDllMappings.PacketOpenAdapter(nifName);

        long hFile = -1;
        if (lpAdapter != null) {
            if (Native.POINTER_SIZE == 4) {
                hFile = lpAdapter.getInt(0);
            } else {
                hFile = lpAdapter.getLong(0);
            }
        }
        if (hFile == -1L) {
            int err = Native.getLastError();
            logger.error("Unable to open the NIF {}, Error Code: {}", nifName, err);
            return null;
        }

        Memory mem = new Memory(NativePacketDllMappings.PACKET_OID_DATA_SIZE);
        mem.clear();
        NativePacketDllMappings.PACKET_OID_DATA oidData = new NativePacketDllMappings.PACKET_OID_DATA(mem);
        oidData.Length = new NativeLong(6L);
        oidData.Oid = new NativeLong(0x01010102L);
        int status = NativePacketDllMappings.PacketRequest(lpAdapter, 0, oidData);
        NativePacketDllMappings.PacketCloseAdapter(lpAdapter);

        if (status == 0) {
            logger.error("Failed to retrieve the link layer address of the NIF: {}", nifName);
            return null;
        } else {
            return MacAddress.getByAddress(oidData.Data);
        }
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder(250);

        sb.append("name: [").append(name).append("] description: [").append(description);

        for (PcapAddress addr : addresses) {
            sb.append("] address: [").append(addr.getAddress());
        }

        for (LinkLayerAddress addr : linkLayerAddresses) {
            sb.append("] link layer address: [").append(addr.getAddress());
        }

        sb.append("] loopBack: [").append(loopBack).append("]");
        sb.append("] up: [").append(up).append("]");
        sb.append("] running: [").append(running).append("]");
        sb.append("] local: [").append(local).append("]");

        return sb.toString();
    }

    @Override
    public int hashCode() {
        final int prime = 31;
        int result = 1;
        result = prime * result + addresses.hashCode();
        result = prime * result + ((description == null) ? 0 : description.hashCode());
        result = prime * result + linkLayerAddresses.hashCode();
        result = prime * result + (local ? 1231 : 1237);
        result = prime * result + (loopBack ? 1231 : 1237);
        result = prime * result + (up ? 1231 : 1237);
        result = prime * result + (running ? 1231 : 1237);
        result = prime * result + name.hashCode();
        return result;
    }

    @Override
    public boolean equals(Object obj) {
        if (this == obj) {
            return true;
        }
        if (obj == null) {
            return false;
        }
        if (!(obj instanceof EditPcapNetworkInterface)) {
            return false;
        }
        EditPcapNetworkInterface other = (EditPcapNetworkInterface) obj;
        if (!addresses.equals(other.addresses)) {
            return false;
        }
        if (description == null) {
            if (other.description != null) {
                return false;
            }
        } else if (!description.equals(other.description)) {
            return false;
        }
        if (!linkLayerAddresses.equals(other.linkLayerAddresses)) {
            return false;
        }
        if (local != other.local) {
            return false;
        }
        if (loopBack != other.loopBack) {
            return false;
        }
        if (up != other.up) {
            return false;
        }
        if (running != other.running) {
            return false;
        }
        if (!name.equals(other.name)) {
            return false;
        }
        return true;
    }
    public static EditPcapNetworkInterface getDevByName(String name) throws PcapNativeException {
        if (name == null) {
            StringBuilder sb = new StringBuilder();
            sb.append("name: ").append(name);
            throw new NullPointerException(sb.toString());
        }

        List<EditPcapNetworkInterface> allDevs = findAllDevs();
        for (EditPcapNetworkInterface pif : allDevs) {
            if (pif.getName().equals(name)) {
                return pif;
            }
        }

        return null;
    }
    private static final Object lock = new Object();
    public static List<EditPcapNetworkInterface> findAllDevs() throws PcapNativeException {
        PointerByReference alldevsPP = new PointerByReference();
        NativeMappings.PcapErrbuf errbuf = new NativeMappings.PcapErrbuf();

        List<EditPcapNetworkInterface> ifList = new ArrayList<>();
        synchronized (lock) {
            int rc = NativeMappings.pcap_findalldevs(alldevsPP, errbuf);
            if (rc != 0) {
                StringBuilder sb = new StringBuilder(50);
                sb.append("Return code: ").append(rc).append(", Message: ").append(errbuf);
                throw new PcapNativeException(sb.toString(), rc);
            }
            if (errbuf.length() != 0) {
                logger.warn("{}", errbuf);
            }

            Pointer alldevsp = alldevsPP.getValue();
            if (alldevsp == null) {
                logger.info("No NIF was found.");
                return Collections.<EditPcapNetworkInterface>emptyList();
            }

            NativeMappings.pcap_if pcapIf = new NativeMappings.pcap_if(alldevsp);

            for (NativeMappings.pcap_if pif = pcapIf; pif != null; pif = pif.next) {
                ifList.add(EditPcapNetworkInterface.newInstance(pif, true));
            }

            NativeMappings.pcap_freealldevs(pcapIf.getPointer());
        }

        logger.info("{} NIF(s) found.", ifList.size());
        return ifList;
    }
}
