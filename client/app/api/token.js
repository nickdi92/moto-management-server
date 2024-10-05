export function IsTokenExpired(token) {
    let tokenData = _parseJWT(token);
    console.debug("TOKENDATA: ", tokenData);
    return false;
}

function _parseJWT(token) {
    if (!token) { return; }
    const base64Url = token.split('.')[1];
    const base64 = base64Url.replace('-', '+').replace('_', '/');
    return JSON.parse(window.atob(base64));
}