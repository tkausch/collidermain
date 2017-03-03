// Copyright (c) 2017 Thomas Kausch. All Rights Reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.

package main

import (
  "github.com/webrtc/apprtc/src/collider/collider"
  "flag"
  "log"
  "os"
  "strconv"
)

var tls = flag.Bool("tls", false, "whether TLS is used")


func main() {
  flag.Parse()

  var port string
  if port = os.Getenv("PORT"); len(port) == 0 {
    port = "8090"
  }

  var roomSrv string
  if roomSrv = os.Getenv("ROOM_SRV"); len(roomSrv) == 0 {
    roomSrv = "https://appr.tc/"
  }


  log.Printf("Starting collider: tls = %t, port = %s, room-server=%s", *tls, port, roomSrv)


  portValue, _ := strconv.Atoi(port)  

  c := collider.NewCollider(roomSrv)
  c.Run(portValue, *tls)
}
