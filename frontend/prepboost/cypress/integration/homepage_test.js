describe("renders the home page", () => {
    it("renders correctly", () => {
      cy.visit("/");
      cy.get("#container").should("exist");
      cy.findAllByText("Login").should("exist");
      cy.findAllByText("Year wise Questions").should("exist");
      cy.contains("Welcome to PREPBOOST").should("exist");
      });

      it("routes to a correct pages", () => {
           cy.findAllByText("Login").click();
           cy.url().should("include", "login");
           cy.findAllByText("Year wise Questions").click();
           cy.url().should("include", "years");
           cy.findAllByText("Topic wise Questions").click();
           cy.url().should("include", "topics");
           cy.findAllByText("Contact").click();
           cy.url().should("include", "contact");
           cy.contains("issues").should("exist");
           cy.findAllByText("Bookmarks").click();
           cy.url().should("include", "bookmarks");
         });
});