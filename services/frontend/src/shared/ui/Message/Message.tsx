export const Message: React.FC<{ title: string; className?: string }> = ({
  title,
  className,
}) => {
  return (
    <div
      className={
        ` block rounded-20 border-accent border-solid border p-2 pt-2 pb-2 pl-4 pr-5 font-light text-sm h-fit` +
        className
      }
    >
      {title}
    </div>
  );
};
