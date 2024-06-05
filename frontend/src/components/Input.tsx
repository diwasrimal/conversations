import { ComponentPropsWithoutRef, forwardRef } from "react";

type InputProps = ComponentPropsWithoutRef<"input">;

type LabeledInputProps = InputProps & {
    id: string;
    label: string;
};

export const Input = forwardRef<HTMLInputElement, InputProps>((props, ref) => {
    return (
        <input
            className="px-2 py-1 bg-[#fafaf5] outline-none"
            ref={ref}
            {...props}
        />
    );
});

export const LabeledInput = forwardRef<HTMLInputElement, LabeledInputProps>(
    ({ label, ...inputProps }, ref) => {
        return (
            <div className="flex flex-col">
                <label htmlFor={inputProps.id}>{label}</label>
                <Input ref={ref} {...inputProps} />
            </div>
        );
    },
);
