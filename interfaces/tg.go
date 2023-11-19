// @tg version=v1.78.10
// @tg http-prefix=api
// @tg security=`bearer`
// @tg title=`Example microservices`
// @tg description=`User Service example`
// @tg packageJSON=`github.com/seniorGolang/json`
// @tg servers=`http://localhost:9000;common`
//
//go:generate tg client -go --services . --ifaces User,Balance --outPath ../clients/servies
//go:generate tg transport --services . --ifaces User,Balance --out ../internal/transport --outSwagger ../../api/user-openapi.json
//go:generate goimports -l -w ../internal/transport ../clients/servies
package interfaces
