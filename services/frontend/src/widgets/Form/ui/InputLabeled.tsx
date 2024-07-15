export const InputLabeled: React.FC<{
  label: string;
  placeholder: string;
  type: React.HTMLInputTypeAttribute | undefined;
  name: string;
  onChange: React.ChangeEventHandler<HTMLInputElement>;
  isIputValueValid: boolean;
}> = ({ label, placeholder, type, name, onChange, isIputValueValid }) => {
  return (
    <div className="grow  even:max-w-[50%] even:min-w-[50%] xl:even:max-w-[35%] xl:even:min-w-[35%] relative">
      <label htmlFor={name} className="text-white text-lg sm:text-xl">
        {label}
      </label>
      <input
        type={type}
        name={name}
        id={name}
        onChange={onChange}
        placeholder={placeholder}
        minLength={2}
        required
        className={` border border-solid border-white rounded-[30px] bg-accent 
          placeholder:text-white placeholder:text-sm sm:placeholder:text-base placeholder:font-light text-white px-5 py-2 w-full mt-2 
          ${isIputValueValid ? "" : "bg-white/30"}
          `}
      />
      <div
        className={`message mt-2 text-white bg-transparent w-full text-sm sm:text-base ${
          isIputValueValid ? "hidden" : "block"
        }`}
      >
        {name === "phone" ? "Введите номер телефона в формате 79995554433" : ""}
        {name === "name"
          ? "Введите имя, содержащее только буквы, пробелы и дефисы."
          : ""}
      </div>
    </div>
  );
};
