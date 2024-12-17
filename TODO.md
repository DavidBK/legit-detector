  - events strings enums
  - const when possible

- Make dispatcher a singleton?
- Error handling
- Error wraps
- Add cleanup function to repo lifetime
- Get secret from secret manager (could be different secrets per client)
- Containerize the ngrok
- Tests
- Use structured logging
- Unmarshal using library for example [json-iter](github.com/json-iterator/go)
- Use [go sdk](https://github.com/octokit/go-sdk/) for github webhook api
(or generate one using the openapi) or using library for example [githubevents](https://github.com/cbrgm/githubevents)
