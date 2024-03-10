# AA6_homework

[ES and data model on miro board](https://miro.com/app/board/uXjVNqVECdA=/?share_link_id=137383829753)

### Testing

Current version assumes some pre-requirements:
 - postgres `postgres:password@localhost:5432`
 - created DBs `auth`, `tasks`, `billing` and `analytics`
 - run docker-compose from root
 - ports of services hard-coded to 8091, 8092, 8093 and 8094

API can be tested by requests from `examples.http` file.