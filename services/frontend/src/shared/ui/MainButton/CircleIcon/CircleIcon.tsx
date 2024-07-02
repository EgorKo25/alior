import styles from "../MainButton.module.scss";

export const CircleIcon: React.FC<{ color: string }> = ({ color }) => (
  <div className={styles.circle_container}>
    <div style={{ backgroundColor: color }} className={styles.circle}></div>
  </div>
);
