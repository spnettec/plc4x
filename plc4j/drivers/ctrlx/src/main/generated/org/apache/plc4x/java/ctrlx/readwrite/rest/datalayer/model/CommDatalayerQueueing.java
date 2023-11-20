/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * ctrlX CORE - Data Layer API
 * This is the base API for the ctrlX Data Layer access on ctrlX CORE devices <ul> <li>Click 'Authorize' to open the 'Available authorizations' dialog.</li> <li>Enter 'username' and 'password'. The 'Client credentials location' selector together with the 'client_id' and 'client_secret' fields as well as the 'Bearer' section can be ignored.</li> <li>Click 'Authorize' and then 'Close' to close the 'Available authorizations' dialog.</li> <li>Try out those GET, PUT, ... operations you're interested in.</li> </ul>
 *
 * The version of the OpenAPI document: 2.1.0
 * Contact: support@boschrexroth.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.model;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonPropertyOrder;
import com.fasterxml.jackson.annotation.JsonTypeName;

import java.io.UnsupportedEncodingException;
import java.net.URLEncoder;
import java.util.Objects;
import java.util.StringJoiner;

/**
 * CommDatalayerQueueing
 */
@JsonPropertyOrder({
  CommDatalayerQueueing.JSON_PROPERTY_QUEUE_SIZE,
  CommDatalayerQueueing.JSON_PROPERTY_BEHAVIOUR
})
@JsonTypeName("comm_datalayer_Queueing")
@jakarta.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-11-18T13:34:36.056861+01:00[Europe/Berlin]")
public class CommDatalayerQueueing {
  public static final String JSON_PROPERTY_QUEUE_SIZE = "queueSize";
  private Integer queueSize = 10;

  public static final String JSON_PROPERTY_BEHAVIOUR = "behaviour";
  private CommDatalayerQueueBehaviour behaviour;

  public CommDatalayerQueueing() {
  }

  public CommDatalayerQueueing queueSize(Integer queueSize) {
    
    this.queueSize = queueSize;
    return this;
  }

   /**
   * size of buffer
   * minimum: 0
   * maximum: 4294967295
   * @return queueSize
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_QUEUE_SIZE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public Integer getQueueSize() {
    return queueSize;
  }


  @JsonProperty(JSON_PROPERTY_QUEUE_SIZE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setQueueSize(Integer queueSize) {
    this.queueSize = queueSize;
  }


  public CommDatalayerQueueing behaviour(CommDatalayerQueueBehaviour behaviour) {
    
    this.behaviour = behaviour;
    return this;
  }

   /**
   * Get behaviour
   * @return behaviour
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_BEHAVIOUR)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public CommDatalayerQueueBehaviour getBehaviour() {
    return behaviour;
  }


  @JsonProperty(JSON_PROPERTY_BEHAVIOUR)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setBehaviour(CommDatalayerQueueBehaviour behaviour) {
    this.behaviour = behaviour;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    CommDatalayerQueueing commDatalayerQueueing = (CommDatalayerQueueing) o;
    return Objects.equals(this.queueSize, commDatalayerQueueing.queueSize) &&
        Objects.equals(this.behaviour, commDatalayerQueueing.behaviour);
  }

  @Override
  public int hashCode() {
    return Objects.hash(queueSize, behaviour);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class CommDatalayerQueueing {\n");
    sb.append("    queueSize: ").append(toIndentedString(queueSize)).append("\n");
    sb.append("    behaviour: ").append(toIndentedString(behaviour)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

  /**
   * Convert the instance into URL query string.
   *
   * @return URL query string
   */
  public String toUrlQueryString() {
    return toUrlQueryString(null);
  }

  /**
   * Convert the instance into URL query string.
   *
   * @param prefix prefix of the query string
   * @return URL query string
   */
  public String toUrlQueryString(String prefix) {
    String suffix = "";
    String containerSuffix = "";
    String containerPrefix = "";
    if (prefix == null) {
      // style=form, explode=true, e.g. /pet?name=cat&type=manx
      prefix = "";
    } else {
      // deepObject style e.g. /pet?id[name]=cat&id[type]=manx
      prefix = prefix + "[";
      suffix = "]";
      containerSuffix = "]";
      containerPrefix = "[";
    }

    StringJoiner joiner = new StringJoiner("&");

    // add `queueSize` to the URL query string
    if (getQueueSize() != null) {
      try {
        joiner.add(String.format("%squeueSize%s=%s", prefix, suffix, URLEncoder.encode(String.valueOf(getQueueSize()), "UTF-8").replaceAll("\\+", "%20")));
      } catch (UnsupportedEncodingException e) {
        // Should never happen, UTF-8 is always supported
        throw new RuntimeException(e);
      }
    }

    // add `behaviour` to the URL query string
    if (getBehaviour() != null) {
      try {
        joiner.add(String.format("%sbehaviour%s=%s", prefix, suffix, URLEncoder.encode(String.valueOf(getBehaviour()), "UTF-8").replaceAll("\\+", "%20")));
      } catch (UnsupportedEncodingException e) {
        // Should never happen, UTF-8 is always supported
        throw new RuntimeException(e);
      }
    }

    return joiner.toString();
  }

}

