describe("renders the home page", () => {
    it("renders correctly", () => {
      cy.visit("/");
      cy.get("#container").should("exist");
      cy.findAllByText("Login").should("exist");
      cy.findAllByText("Year wise Questions").should("exist");
      cy.contains("Welcome to PREPBOOST").should("exist");
      });
});