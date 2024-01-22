This guide describes how to generate a new .proto (only needed when there is a new OpenPrio specification).

1. Download newest specification https://github.com/openprio/specification/tree/master.
2. Add ```option go_package = "/openprio_pt_position_data";``` on top of .proto file.
3. Run ```protoc --proto_path=. --go_out=. *.proto```
