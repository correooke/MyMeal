"use client";

import React from "react";
import Grid from "@mui/material/Grid";
import Paper from "@mui/material/Paper"; // For demo purposes, to visualize the layout
import { styled } from "@mui/material/styles";

// Styling the Paper component for visual purposes
const Item = styled(Paper)(({ theme }) => ({
  padding: theme.spacing(1),
  textAlign: "center",
  color: theme.palette.text.secondary,
}));

const SideCard = () => (
  <Item style={{ backgroundColor: "red", height: "100%" }}>SideCard</Item>
);
const BusinessForm = () => (
  <Item style={{ backgroundColor: "blue", height: "100%" }}>BusinessForm</Item>
);
const CategoryForm = () => (
  <Item style={{ backgroundColor: "green", height: "100%" }}>CategoryForm</Item>
);

function MyLayout() {
  return (
    <Grid container spacing={0} style={{ height: "100vh" }}>
      <Grid item xs={3}>
        <SideCard />
      </Grid>
      <Grid item xs={9} container direction="column">
        <Grid item xs={3}>
          <BusinessForm />
        </Grid>
        <Grid item xs={9}>
          <CategoryForm />
        </Grid>
      </Grid>
    </Grid>
  );
}

export default MyLayout;
