'use client';

import {useRouter} from "next/navigation";
import {GetUserDataFromLocalStorage, GetUserFullName, GetUsername, IsUserLoggedIn} from "@/app/helpers/userHelper";
import DashboardHeader from "@/app/components/dashboard/layout";
import {getNavigationUrl} from "@/app/helpers/utils";

export default function Dashboard() {
    const router = useRouter();
    if (!IsUserLoggedIn()) {
        router.push("/admin/login");
    } else {
        let userInfo = GetUserDataFromLocalStorage();
        let motorcyclesNode = userInfo?.motorcycles;
        let motorcyclesCounter = motorcyclesNode && motorcyclesNode.length > 0 ? motorcyclesNode.length : 0;
        return (
            <>
                <div className="min-h-screen">
                    {<DashboardHeader title={"Dashboard"}/>}
                    <main className="bg-white shadow min-h-screen">
                        <div className="mx-auto max-w-7xl px-4 py-6 sm:px-6 lg:px-8 ">
                            <div className="max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700">
                                <a href="#">
                                    <h5 className="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
                                        {`You have ${motorcyclesCounter} Motorcycles`}
                                    </h5>
                                </a>
                                <p className="mb-3 font-normal text-gray-700 dark:text-gray-400">
                                    Go to the Motorcycles section to manage your data
                                </p>
                                <a href={getNavigationUrl("Motorcycles")} className="inline-flex items-center px-3 py-2 text-sm font-medium
                                text-center text-white bg-blue-700 rounded-lg hover:bg-blue-800 focus:ring-4
                                focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700
                                dark:focus:ring-blue-800">
                                    Go to
                                    <svg className="rtl:rotate-180 w-3.5 h-3.5 ms-2" aria-hidden="true"
                                         xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 10">
                                        <path stroke="currentColor"
                                              d="M1 5h12m0 0L9 1m4 4L9 9"/>
                                    </svg>
                                </a>
                            </div>
                        
                        </div>
                    </main>
                </div>
            </>
        )
    }
}
