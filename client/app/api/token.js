import {jwtDecode} from "jwt-decode";
import {getCookie} from "cookies-next";

export function IsTokenExpired() {
    let token = getCookie("bearer_token");
    if (!token) return true;
    let tokenData = jwtDecode(token);
    let now = Date.now() / 1000;
    return tokenData.exp < now;
}
