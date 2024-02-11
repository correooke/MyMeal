import { TextField } from "@mui/material";
import { InputProps } from "../../Shared/StringProps";

const ProductInput = ({ value, onChange }: InputProps) => (
  <TextField
    name={"product"}
    value={value}
    label={"Producto"}
    onChange={(event) => onChange(event.target.value, event.target.name)}
    fullWidth
    margin="normal"
  />
);

export default ProductInput;
