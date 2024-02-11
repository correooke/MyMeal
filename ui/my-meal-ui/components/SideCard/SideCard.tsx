"use client";

import React from "react";
import { Box, CardMedia } from "@mui/material";
import { useTheme } from "@mui/material/styles";

const SideCard = () => {
  const theme = useTheme();
  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        width: "100%",
      }}
    >
      <CardMedia
        component="img"
        src="/resources/images/logo.png"
        sx={{
          height: theme.header.height,
          border: "0",
          margin: "0px",
          objectFit: "cover",
          objectPosition: "center",
        }}
      ></CardMedia>
      <Box
        sx={{
          width: "100%",
          height: "100%",
          border: "0",
          margin: "0px",
          backgroundColor: theme.palette.primary.main,
        }}
      ></Box>
    </Box>
  );
};

export default SideCard;
