export function GetUserFullName() {
    let user = GetUser();
    if (user) {
        return user?.name && user?.lastname ?
            user.name + " " + user.lastname :
            user.username;
    }
    return "";
}

export function GetUser() {
    return JSON.parse(localStorage.getItem("user_data")) ?? null;
}