import "./Button.css";

export default function Button({ children, style, ...rest }) {
    // Used in ./Button.css
    const vars = {
        "--width": style?.width || "130px",
        "--height": style?.height || "40px",
        "--bg-color": style?.backgroundColor || "var(--grey)"
    };

    return (
        <button className="button" style={vars} {...rest}>
            {children}
        </button>
    );
}
