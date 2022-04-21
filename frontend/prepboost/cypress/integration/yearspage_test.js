describe("renders the years page", () => {
    it("renders correctly", () => {
      cy.visit("/years");
      cy.contains("different years").should("exist");
      cy.findAllByText("2021").should("exist");
      cy.findAllByText("2020").should("exist");
      cy.findAllByText("2019").should("exist");
      cy.findAllByText("2011").should("exist");
      });

      it("routes to a correct pages", () => {
           cy.findAllByText("2020").click();
           cy.url().should("include", "2020questions");
           cy.contains("scored").should("exist");
           cy.contains("Question").should("exist");
         });
});