import { forwardRef } from "react";
import "../styles/InputFields.css";

const LabeledInputField = forwardRef(({ label, ...inputProps }, ref) => {
    const id = `${label.replace(" ", "-")}-id`;
    return (
        <div className="labeled-input-field-container">
            <label htmlFor={id}>{label}</label>
            <InputField id={id} ref={ref} {...inputProps} />
        </div>
    );
});

const InputField = forwardRef((inputProps, ref) => {
    return <input className="input" ref={ref} {...inputProps} />;
});

export { InputField, LabeledInputField };
