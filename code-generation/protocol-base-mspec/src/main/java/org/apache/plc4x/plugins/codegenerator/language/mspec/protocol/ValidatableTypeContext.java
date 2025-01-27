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
package org.apache.plc4x.plugins.codegenerator.language.mspec.protocol;

import org.apache.plc4x.plugins.codegenerator.protocol.TypeContext;
import org.apache.plc4x.plugins.codegenerator.types.exceptions.GenerationException;

public interface ValidatableTypeContext extends TypeContext {

    /**
     * validates the {@link TypeContext}
     *
     * @throws GenerationException if {@link TypeContext}
     */
    default void validate() throws GenerationException {
        // TODO: check that we have at least of parsed type
        if (getUnresolvedTypeReferences().size() > 0) {
            throw new GenerationException("Unresolved types left: " + getUnresolvedTypeReferences());
        }
    }
}
