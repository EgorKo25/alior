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
      className={clsx(styles.button, className) + ` md:w-20 md:h-20`}
      type={type}
      onClick={onClick}
    >
      <TgIcon color="#FFFFFF" className=" w-10 h-10 md:w-12 md:h-12" />
    </button>
  );
};
