describe('empty spec', () => {
    it('passes', () => {
      cy.visit('http://localhost:3000/')
      cy.get('.cart-link').first().click()
      cy.url().should('include', '/cart')
      
      cy.get('.home-link').first().click()
      cy.url().should('include', '/')
    })
  })