package middlewares

// A map of routes and the corresponding roles that can use them
var Routes = map[string][]string {
	"GET /inventories/software": {
		"admin",
		"moderator",
		"user",
	},
	"GET /inventories/software/:id": {
		"admin",
		"moderator",
		"user",
	},
	"POST /inventories/software": {
		"admin",
		"moderator",
	},
	"PUT /inventories/software/:id": {
		"admin",
		"moderator",
	},
	"DELETE /inventories/software/:id": {
		"admin",
	},
	"GET /inventories/ssh": {
		"admin",
		"moderator",
		"user",
	},
	"GET /inventories/ssh/:id": {
		"admin",
		"moderator",
		"user",
	},
	"POST /inventories/ssh": {
		"admin",
		"moderator",
	},
	"PUT /inventories/ssh/:id": {
		"admin",
		"moderator",
	},
	"DELETE /inventories/ssh/:id": {
		"admin",
	},
	"GET /users": {
		"admin",
		"moderator",
	},
	"GET /users/:id": {
		"admin",
		"moderator",
	},
	"POST /users": {
		"admin",
	},
	"PUT /users/:id": {
		"admin",
	},
	"PATCH /users/:id": {
		"admin",
	},
	"DELETE /users/:id": {
		"admin",
	},
	"GET /inventories/servers": {
		"admin",
	},
	"GET /inventories/servers/:id": {
		"admin",
	},
	"GET /inventories/serversSSH/": {
		"admin",
		"moderator",
		"user",
	},
	"GET /inventories/serversSoftware/:id": {
		"admin",
		"moderator",
		"user",
	},
	"GET /serversByUser": {
		"admin",
		"moderator",
		"user",
	},
	"POST /inventories/servers": {
		"admin",
		"moderator",
	},
}

