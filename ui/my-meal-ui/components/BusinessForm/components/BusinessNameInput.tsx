import { TextField } from "@mui/material";
import { InputProps } from "../../Shared/StringProps";

const BusinessNameInput = ({ value, onChange }: InputProps) => (
  <TextField
    name={"business-name"}
    value={value}
    label={"Nombre del negocio"}
    onChange={(event) => onChange(event.target.value, event.target.name)}
    margin="none"
    sx={{ width: "800px" }}
  />
);

export default BusinessNameInput;
