describe("renders the topics page", () => {
    it("renders correctly", () => {
      cy.visit("/topics");
      cy.contains("different topics").should("exist");
      cy.findAllByText("Economy").should("exist");
      cy.findAllByText("Agriculture").should("exist");
      cy.findAllByText("Environment").should("exist");
      cy.findAllByText("Miscellaneous").should("exist");
      });

      it("routes to a correct pages", () => {
           cy.findAllByText("Miscellaneous").click();
           cy.url().should("include", "miscellaneousquestions");
           cy.contains("scored").should("exist");
          cy.contains("Question").should("exist");
         });

      it("routes to a correct pages", () => {
        cy.visit("/topics");
          cy.findAllByText("Economy").click();
          cy.url().should("include", "economyquestions");
          cy.contains("scored").should("exist");
          cy.contains("Question").should("exist");
        });
});