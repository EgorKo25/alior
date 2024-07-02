import clsx from "clsx";
import styles from "./TgButton.module.scss";
import React from "react";
import { TgIcon } from "./TgIcon";

export const TgButton = ({
  onClick,
  type,
  className,
}: {
  onClick?: () => void;
  type?: React.ButtonHTMLAttributes<HTMLButtonElement>["type"];
  className?: string;
}) => {
  return (
    <button
      className={clsx(styles.button, className)}
      type={type}
      onClick={onClick}
    >
      <TgIcon />
    </button>
  );
};
