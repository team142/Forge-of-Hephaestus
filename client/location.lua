-- Send a location to the server
x,y,z = gps.location()
body = args[0]
http.post("http://192.168.88.161:8080/api/location", body)

