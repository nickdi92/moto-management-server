import {getCookie, setCookie} from "cookies-next";

export function GetAuthenticationHeader() {
    return   "Bearer " + GetBearerToken();
}

export function GetBearerToken() {
    return getCookie("bearer_token")
}

export function UpdateBearerToken(newToken) {
    setCookie("bearer_token", newToken)
}