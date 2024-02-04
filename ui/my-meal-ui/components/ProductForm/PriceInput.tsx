import { TextField } from "@mui/material";
import { NumberProps } from "./NumberProps";

const PriceInput = ({ value, onChange }: NumberProps) => (
  <TextField
    name={"price"}
    value={value}
    label={"Precio"}
    onChange={(event) =>
      onChange(parseInt(event.target.value), event.target.name)
    }
    fullWidth
    margin="normal"
    type="number"
  />
);

export default PriceInput;
