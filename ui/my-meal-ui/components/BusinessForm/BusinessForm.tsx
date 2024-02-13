"use client";

import { FC } from "react";
import { Box } from "@mui/material";
import BusinessImage from "./components/BusinessImage";
import BusinessLogo from "./components/BusinessLogo";
import BusinessNameInput from "./components/BusinessNameInput";
import QRCodeGeneratorButton from "./components/QRCodeGeneratorButton";
import PreviewButton from "./components/PreviewButton";
import CurrencySelector from "./components/CurrencySelector";

const BusinessForm: FC = () => {
  return (
    <Box sx={{ width: "100%", display: "block", height: "220px" }}>
      <BusinessImage />
      <Box
        sx={{
          display: "flex",
          flexDirection: "row",
          alignSelf: "center",
          padding: "10px",
          gap: "10px",
        }}
      >
        <BusinessLogo />
        <BusinessNameInput value={"Las violetas"} onChange={() => {}} />
        <CurrencySelector />
        <Box
          sx={{
            width: "100%",
            display: "flex",
            gap: "10px",
            justifyContent: "flex-end",
            padding: "10px",
          }}
        >
          <QRCodeGeneratorButton />
          <PreviewButton />
        </Box>
      </Box>
    </Box>
  );
};

export default BusinessForm;
