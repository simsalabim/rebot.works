### TODO
- [x] Landing page
- [ ] Database integration: save searched terms and download displayed images.
- [ ] Background job to crawl all images related to recent search terms.

### Deployment
 - build the binary for linux with `env GOOS=linux GOARCH=amd64 go build giphy`
 - scp the binary and html files to the destination server
 - install [daemonize](http://software.clapper.org/daemonize/) on the server
 - `export WEATHER_APP_ID=`cat /home/root/apps/rebot.works/.weather_app_id``
 - run the program on the server with `daemonize -c /home/root/apps/rebot.works -p /home/root/apps/rebot.works/myapp.pid /home/root/apps/rebot.works/giphy`
 - grow up and use anything other than `scp` for deployments :)
