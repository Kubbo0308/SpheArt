export const CONST = {
  API_BASE_PATH: "http://backend:8080",
  TOP: "/",
  CSRF_TOKEN: "/csrf",
  ARTICLES: "/articles",
  SEARCH_ARTICLES_IN_TITLE: "/articles/search?title=",
  AUTH: "/auth",
  SIGN_UP: "/signup",
  SIGN_IN: "/signin",
  SIGN_OUT: "/signout"
}

export const STATUS_CODE = {
  OK: 200,
  CREATED: 201,
  NO_CONTENT: 204,
  BAD_REQUEST: 400,
  UNAUTHORIZED: 401,
  FORBIDEN: 403,
  NOT_FOUND: 404,
  CONFLICT: 409,
  UNPROCESSABLE_ENTITY: 422,
  INTERNAL_SERVER_ERROR: 500,
  ERROR: 999
} as const