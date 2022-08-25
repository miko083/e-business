describe('empty spec', () => {
    it('passes', () => {
      cy.visit('http://localhost:3000/')
      cy.get('.product-add-button').first().click()
      cy.url().should('include', '/')
      
      cy.visit('http://localhost:3000/cart')
      cy.url().should('include', '/cart')

      cy.get('.cart-items').find('.item-in-cart').should('have.length', 1)

      cy.visit('http://localhost:3000/')
      cy.get('.home-link').click()

      cy.url().should('include', '/')
      cy.get('.product-add-button').eq(1).click()

      cy.visit('http://localhost:3000/cart')
      cy.url().should('include', '/cart')
      cy.get('.cart-items').find('.item-in-cart').should('have.length', 2)

      cy.get('.cart-remove-button').first().click()

      cy.get('.cart-items').find('.item-in-cart').should('have.length', 1)


    })
  })