import styles from './styles.module.css'
export default function Logo({
    context
}) {
    let logoStyle = context === "dashboard";
    return (
        <>
            <div className={logoStyle ? styles.logo : styles.logoGuest}>
                <h1>Motor Service</h1>
            </div>
        </>
    )
}