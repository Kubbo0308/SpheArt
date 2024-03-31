import { NextRequest, NextResponse } from "next/server";

export function middleware(request: NextRequest) {
  const cookie = request.cookies.get('token')
  if (cookie === undefined) {
    if (request.nextUrl.pathname === "/bookmark") {
      return NextResponse.redirect(new URL("/", request.url))
    }
  }
}