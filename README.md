# Valkiria

Pet project for self development.

## TODO

in no particular order...

- ~~[ ] Create gui with [systray](https://pkg.go.dev/github.com/getlantern/systray@v1.2.2#section-readme).~~
- ~~[ ] Write DB k8s manifest~~
- [x] Update Dockerfile with db → separatared Dockerfiles for database and application
- [x] Create docker-compose
- [x] Declare table creation in code
- [x] Update request handlers to call movies table
- [x] Refactor db.Dockerfile in favour of docker-compose
- [x] Make valkiria-app wait for valkiria-postgres
- [ ] Create app as helm3 chart
- [ ] Move postgres db to AWS RDS
- [ ] Implement terraform to manage AWS resource
