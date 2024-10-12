import {TailSpin} from "react-loader-spinner";
import React from "react";
import styles from './styles.module.css'

export default function MotoLoader() {
    return <TailSpin
        visible={true}
        height="80"
        width="80"
        color="#4fa94d"
        ariaLabel="tail-spin-loading"
        radius="1"
        wrapperStyle={{}}
        wrapperClass={styles.formLoader}
    />
}