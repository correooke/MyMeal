import React, { useState } from "react";
import {
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  SelectChangeEvent,
} from "@mui/material";

const CurrencySelector = () => {
  const [currency, setCurrency] = useState(() => "pesos");

  const handleChange = (event: SelectChangeEvent<{ value: unknown }>) => {
    setCurrency(event.target.value as string);
  };

  return (
    <FormControl sx={{ width: "350px" }}>
      <InputLabel id="currency-label">Precio expresado en</InputLabel>
      <Select
        labelId="currency-label"
        id="currency-select"
        value={currency as "" | { value: unknown } | undefined}
        onChange={handleChange}
      >
        <MenuItem value="pesos">Pesos</MenuItem>
        <MenuItem value="dollars">Dólares</MenuItem>
      </Select>
    </FormControl>
  );
};

export default CurrencySelector;
