export function classNames(...classes) {
    return classes.filter(Boolean).join(' ')
}

export function getNavigationUrl(navigationName) {
    let url = GetNavigationMenuItems().filter((nav) => {
        if (nav.name === navigationName) {
            return nav.href;
        }
        return "";
    })
    
    return url.length == 1 ? url[0].href : "#";
}

export function GetNavigationMenuItems() {
    return [
        {name: 'Dashboard', href: '/admin/dashboard'},
        {name: 'Motorcycles', href: '/admin/motorcycles'},
        {name: 'Fuels', href: '/admin/motorcycles/fuels'},
        {name: 'Services', href: '/admin/motorcycles/service'},
        {name: 'Accidents', href: '/admin/motorcycles/accidents'},
    ];
}

export function IsCurrentRoute(route) {
    return route === window.location.pathname || window.location.pathname.includes(route);
}

export function GetUserNavigationItems() {
    return [
        {name: 'Your Profile', href: '/admin/user/'},
        {name: 'Settings', href: '#'},
        {name: 'Sign out', href: '#'},
    ]
}