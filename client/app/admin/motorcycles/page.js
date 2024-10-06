"use client";

import DashboardHeader from "@/app/components/dashboard/layout";

export default function Motorcycles() {
    return (
        <>
            <div className="min-h-full">
                {<DashboardHeader title={"Motorcycles"}/>}
            </div>
        </>
    );
}