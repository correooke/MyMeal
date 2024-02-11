"use client";

import { FC } from "react";
import { Box } from "@mui/material";
import BusinessImage from "./components/BusinessImage";
import BusinessLogo from "./components/BusinessLogo";
import BusinessNameInput from "./components/BusinessNameInput";
import QRCodeGeneratorButton from "./components/QRCodeGeneratorButton";
import PreviewButton from "./components/PreviewButton";

const BusinessForm: FC = () => {
  return (
    <Box width={"100%"}>
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
        <QRCodeGeneratorButton />
        <PreviewButton />
      </Box>
    </Box>
  );
};

export default BusinessForm;
