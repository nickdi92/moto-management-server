import {IsTokenExpired} from "@/app/api/token";
import {UpdateBearerToken} from "@/app/api/auth";
import {getCookie, setCookie} from "cookies-next";

export function GetUserFullName() {
    let user = GetUserDataFromLocalStorage();
    if (user) {
        return user?.name && user?.lastname ?
            user.name + " " + user.lastname :
            user.username;
    }
    return "";
}

export function GetUserDataFromLocalStorage() {
    return JSON.parse(localStorage.getItem("user_data")) ?? null;
}

export function UpdateUserDataToLocalStorage(userData) {
    localStorage.setItem("user_data", JSON.stringify(userData));
}

export function IsUserLoggedIn() {
    let isUserLoggedIn = Boolean(localStorage.getItem("is_user_logged_in"));
    return !IsTokenExpired() && isUserLoggedIn;
}

export function SetIsUserLoggedIn(userData) {
    localStorage.setItem("is_user_logged_in", userData?.is_logged_in);
    UpdateBearerToken(userData.token);
}

export function SetUsername(username) {
    if (username) {
        setCookie("username", username)
    }
}

export function GetUsername() {
    return getCookie("username");
}