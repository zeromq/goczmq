A go interface to [CZMQ](http://czmq.zeromq.org)

This requires CZMQ head, and is targetted to be compatible with the next stable release of CZMQ.

Development is currently using CZMQ head compiled against ZeroMQ 4.0.4 Stable.

## Install

  go get github.com/zeromq/goczmq

## Currently Working

* Bare bones zsock functionality.  Sockets can be created and destroyed, and messages can be sent.

## Next Objectives

* Finish support of zsock api, working with sending go byte arrays.

## See Also

Peter Kleiweg's excellent zmq4 library for libzmq: http://github.com/pebbe/zmq4
