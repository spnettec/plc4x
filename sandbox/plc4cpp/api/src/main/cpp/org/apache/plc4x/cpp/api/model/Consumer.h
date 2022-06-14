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

#ifndef _PLC_CONSUMER
#define _PLC_CONSUMER

#include <functional>

namespace org
{
	namespace apache
	{
		namespace plc4x
		{
			namespace cpp
			{
				namespace api
				{
					namespace model
					{						
						/**
						 * helper class that implements java.util.function.Consumer
						 */
						template<typename T> class Consumer
						{
							
							public:									
								template<class consumer>
								Consumer<T>(consumer consume) : consume_(consume) {};
																					
								inline void Accept(T t) { consume_(t); };
								//inline void andthen(consumer c) {c.Accept(T)}; //TODO
							private:
								T innerType;
								std::function<void(T)> consume_;

						};
						
					}
				}
			}
		}
	}
}

#endif