### TODO
- [x] Landing page
- [ ] Database integration: save searched terms and download displayed images.
- [ ] Background job to crawl all images related to recent search terms.
- [x] Grow up and use ~~anything other than `scp`~~ :elephant: [Gradle](https://gradle.org) for deployments

### Deployment
Deployment scripts rely on `user` and `hostname` variables, so set them up first: `export user=mylogin && export hostname=127.0.0.1`, and `./gradlew deploy` second :metal:.
