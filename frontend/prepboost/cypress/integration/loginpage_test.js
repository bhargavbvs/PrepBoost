describe("renders the home page", () => {
    it("renders correctly", () => {
      cy.visit("/login");
      cy.findAllByText("Username").should("exist");
      cy.findAllByText("Password").should("exist");
      cy.findAllByText("Remember me").should("exist");
      cy.findAllByText("Forgot password ?").should("exist");
      cy.findAllByText("Sign Up").should("exist");
      });

      it("routes to a correct pages", () => {
           cy.findAllByText("Sign Up").click();
           cy.url().should("include", "signup");
         });
});