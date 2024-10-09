"use client"

import {useRouter} from "next/navigation";
import {GetUserDataFromLocalStorage, IsUserLoggedIn} from "@/app/helpers/userHelper";
import DashboardHeader from "@/app/components/dashboard/layout";

export default function User() {
    const router = useRouter();
    if (!IsUserLoggedIn()) {
        router.push("/admin/login");
    } else {
        let userInfo = GetUserDataFromLocalStorage();
        return (
            <>
                <div className="min-h-full bg-white shadow">
                    {<DashboardHeader title={userInfo.name + " " + userInfo.lastname}/>}
                    
                    <main className="bg-white shadow min-h-screen">
                        <div className="mx-auto max-w-7xl px-4 py-6 sm:px-6 lg:px-8 ">
                            <form>
                                <div className="space-y-12">
                                    <div className="border-b border-gray-900/10 pb-12">
                                        <h2 className="text-base font-semibold leading-7 text-gray-900">
                                            Personal Information
                                        </h2>
                                        <div className="mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-9 mb-10">
                                            <div className="sm:col-span-3">
                                                <label htmlFor="firstname"
                                                       className="block text-sm font-medium leading-6 text-gray-900">
                                                    First name
                                                </label>
                                                <div className="mt-2">
                                                    <input
                                                        id="firstname"
                                                        name="firstname"
                                                        value={userInfo.name}
                                                        type="text"
                                                        autoComplete="given-name"
                                                        className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                                    />
                                                </div>
                                            </div>
                                            
                                            <div className="sm:col-span-3">
                                                <label htmlFor="lastname"
                                                       className="block text-sm font-medium leading-6 text-gray-900">
                                                    Last name
                                                </label>
                                                <div className="mt-2">
                                                    <input
                                                        id="lastname"
                                                        name="lastname"
                                                        value={userInfo.lastname}
                                                        type="text"
                                                        autoComplete="family-name"
                                                        className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                                    />
                                                </div>
                                            </div>
                                            
                                            <div className="sm:col-span-3">
                                                <label htmlFor="email"
                                                       className="block text-sm font-medium leading-6 text-gray-900">
                                                    Email address
                                                </label>
                                                <div className="mt-2">
                                                    <input
                                                        id="email"
                                                        name="email"
                                                        type="email"
                                                        value={userInfo.email}
                                                        autoComplete="email"
                                                        className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                                    />
                                                </div>
                                            </div>
                                        </div>
                                        <h2 className="text-base font-semibold leading-7 text-gray-900">
                                            Address Information
                                        </h2>
                                        <div className="mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-9 mb-10">
                                            <div className="sm:col-span-3">
                                                <label htmlFor="street"
                                                       className="block text-sm font-medium leading-6 text-gray-900">
                                                    Street address
                                                </label>
                                                <div className="mt-2">
                                                    <input
                                                        id="street"
                                                        name="street"
                                                        value={userInfo?.address?.street}
                                                        type="text"
                                                        autoComplete="street"
                                                        className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                                    />
                                                </div>
                                            </div>
                                            
                                            <div className="sm:col-span-3">
                                                <label htmlFor="city"
                                                       className="block text-sm font-medium leading-6 text-gray-900">
                                                    City
                                                </label>
                                                <div className="mt-2">
                                                    <input
                                                        id="city"
                                                        name="city"
                                                        type="text"
                                                        value={userInfo?.address?.city}
                                                        autoComplete="address-level2"
                                                        className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                                    />
                                                </div>
                                            </div>
                                            
                                            <div className="sm:col-span-3">
                                                <label htmlFor="province"
                                                       className="block text-sm font-medium leading-6 text-gray-900">
                                                    Province
                                                </label>
                                                <div className="mt-2">
                                                    <input
                                                        id="province"
                                                        name="province"
                                                        value={userInfo?.address?.province}
                                                        type="text"
                                                        autoComplete="address-level1"
                                                        className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                                    />
                                                </div>
                                            </div>
                                            
                                            <div className="sm:col-span-3">
                                                <label htmlFor="zip_code"
                                                       className="block text-sm font-medium leading-6 text-gray-900">
                                                    ZIP / Postal code
                                                </label>
                                                <div className="mt-2">
                                                    <input
                                                        id="zip_code"
                                                        name="zip_code"
                                                        value={userInfo?.address?.zip_code}
                                                        type="text"
                                                        autoComplete="postal-code"
                                                        className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                                    />
                                                </div>
                                            </div>
                                            <div className="sm:col-span-3">
                                                <label htmlFor="state"
                                                       className="block text-sm font-medium leading-6 text-gray-900">
                                                    Country
                                                </label>
                                                <div className="mt-2">
                                                    <input
                                                        id="state"
                                                        name="state"
                                                        value={userInfo?.address?.state}
                                                        type="text"
                                                        autoComplete="postal-code"
                                                        className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                                    />
                                                </div>
                                            </div>
                                        </div>
                                        
                                        <h2 className="text-base font-semibold leading-7 text-gray-900">
                                            Registry Information
                                        </h2>
                                        <div className="mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-9 mb-10">
                                            <div className="sm:col-span-3">
                                                <label htmlFor="fiscal_code"
                                                       className="block text-sm font-medium leading-6 text-gray-900">
                                                    Fiscal code
                                                </label>
                                                <div className="mt-2">
                                                    <input
                                                        id="fiscal_code"
                                                        name="fiscal_code"
                                                        value={userInfo?.registry?.fiscal_code}
                                                        type="text"
                                                        autoComplete="street-address"
                                                        className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                                    />
                                                </div>
                                            </div>
                                            
                                            <div className="sm:col-span-3">
                                                <label htmlFor="dob"
                                                       className="block text-sm font-medium leading-6 text-gray-900">
                                                    Date Of Birth
                                                </label>
                                                <div className="mt-2">
                                                    <input
                                                        id="dob"
                                                        name="dob"
                                                        type="text"
                                                        value={userInfo?.registry?.dob}
                                                        autoComplete="address-level2"
                                                        className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                                    />
                                                </div>
                                            </div>
                                        </div>
                                        
                                    </div>
                                </div>
                                
                                <div className="mt-6 flex items-center justify-end gap-x-6">
                                    <button type="button" className="text-sm font-semibold leading-6 text-gray-900">
                                        Cancel
                                    </button>
                                    <button
                                        type="submit"
                                        className="rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                                    >
                                        Save
                                    </button>
                                </div>
                            </form>
                        </div>
                    </main>
                </div>
            </>
        )
    }
}