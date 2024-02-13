import BusinessForm from "@/components/BusinessForm/BusinessForm";
import CategoryForm from "@/components/CategoryForm/CategoryForm";
import SideCard from "@/components/SideCard/SideCard";
import { Box, Container } from "@mui/material";
import Grid from "@mui/material/Grid";

export default function Home() {
  return (
    <Container maxWidth={false} disableGutters>
      <Grid
        container
        spacing={0}
        style={{ height: "100vh" }}
        sx={{ padding: "0px" }}
      >
        <Grid item xs={2} container sx={{ padding: "0px" }}>
          <SideCard />
        </Grid>
        <Grid item xs={10} container spacing={0}>
          <Box sx={{ width: "100%", height: "25vh" }}>
            <BusinessForm />
          </Box>
          <Box
            sx={{
              display: "flex",
              width: "100%",
              height: "75vh",
              justifyContent: "center",
              alignItems: "top",
              overflow: "auto",
            }}
          >
            <CategoryForm />
          </Box>
        </Grid>
      </Grid>
    </Container>
  );
}
