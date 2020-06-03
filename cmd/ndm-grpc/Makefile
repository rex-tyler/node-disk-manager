.PHONY: protos

protos:
	 protoc -I protos/ protos/ndm.proto --go_out=plugins=grpc:protos/ndm

pack:
	pack build harshthakur/ndm-grpc --builder=cloudfoundry/cnb:tiny --publish