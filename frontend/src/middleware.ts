import { NextRequest, NextResponse } from "next/server";
import { CONST } from "./const";

export function middleware(request: NextRequest) {
  const cookie = request.cookies.get('token')
  if (cookie === undefined) {
    if (request.nextUrl.pathname === CONST.BOOKMARK) {
      return NextResponse.redirect(new URL(`${CONST.AUTH}${CONST.SIGN_IN}`, request.url))
    }

    if (request.nextUrl.pathname === `${CONST.AUTH}${CONST.SIGN_OUT}`) {
      return NextResponse.redirect(new URL(CONST.TOP, request.url))
    }
  } else {
    if (request.nextUrl.pathname === `${CONST.AUTH}${CONST.SIGN_IN}`) {
      return NextResponse.redirect(new URL(CONST.TOP, request.url))
    }
  }
}
