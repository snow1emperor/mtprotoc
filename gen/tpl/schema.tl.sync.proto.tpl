/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'codegen_proto.py'
 *
 *  Copyright (c) 2017, https://github.com/snow1emperor
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

syntax = "proto3";

package mtproto;

option java_package = "com.snow1emperor.engine.mtproto";
option java_outer_classname = "MTProto";
option optimize_for = CODE_SIZE;

option go_package = "mtproto";

import "schema.tl.crc32.proto";
import "schema.tl.core_types.proto";

{{ range .BaseTypeList }}
///////////////////////////////////////////////////////////////////////////////
// {{.Name}} <--
{{ range .SubMessageList }}//  + TL_{{.Name}}
{{end}}//
message {{.Name}}_Data {
{{range .ParamList}}    {{.Type}} {{.Name}} = {{.Index}};
{{end}}}

message {{.Name}} {
    TLConstructor constructor = 1;
    {{.Name}}_Data data2 = 2;
}

{{ range .SubMessageList }}// {{.Line}}
message TL_{{.Name}} {
    {{.ResType}}_Data data2 = 2;
}

{{end}}{{end}}
