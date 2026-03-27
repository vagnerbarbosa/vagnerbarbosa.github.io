/**
 * Componente Accordion
 * @module Accordion
 * @requires App.Utils
 */

(function() {
  'use strict';

  window.App = window.App || {};

  /**
   * Gerenciador de accordions
   * @namespace App.Accordion
   */
  App.Accordion = {
    /**
     * Seletor CSS dos accordions
     * @constant {string}
     */
    ACCORDION_SELECTOR: '.accordion',

    /**
     * Seletor CSS dos headers
     * @constant {string}
     */
    HEADER_SELECTOR: '.accordion-header',

    /**
     * Classe CSS para accordion expandido
     * @constant {string}
     */
    EXPANDED_CLASS: 'expanded',

    /**
     * Atributo ARIA para expansão
     * @constant {string}
     */
    ARIA_EXPANDED: 'aria-expanded',

    /**
     * Cache de elementos
     * @private
     */
    _elements: {
      accordions: null,
      headers: null
    },

    /**
     * Inicializa o módulo de accordions
     * @returns {boolean}
     */
    init: function() {
      try {
        this._cacheElements();

        if (!this._validateElements()) {
          console.warn('Accordion: No accordion elements found');
          return false;
        }

        this._bindEvents();

        return true;
      } catch (error) {
        console.error('Accordion initialization failed:', error);
        return false;
      }
    },

    /**
     * Cache de elementos
     * @private
     */
    _cacheElements: function() {
      this._elements.headers = document.querySelectorAll(this.HEADER_SELECTOR);
      this._elements.accordions = document.querySelectorAll(this.ACCORDION_SELECTOR);
    },

    /**
     * Validação de elementos
     * @private
     * @returns {boolean}
     */
    _validateElements: function() {
      return this._elements.headers &&
             this._elements.headers.length > 0;
    },

    /**
     * Vincula eventos
     * @private
     */
    _bindEvents: function() {
      const self = this;

      // Usa event delegation em vez de listeners individuais
      document.addEventListener('click', function(event) {
        const header = event.target.closest(self.HEADER_SELECTOR);

        if (header) {
          event.preventDefault();
          self._handleHeaderClick(header);
        }
      });

      // Suporte a teclado
      document.addEventListener('keydown', function(event) {
        if (event.key === 'Enter' || event.key === ' ') {
          const header = event.target.closest(self.HEADER_SELECTOR);

          if (header) {
            event.preventDefault();
            self._handleHeaderClick(header);
          }
        }
      });
    },

    /**
     * Handler de clique no header
     * @private
     * @param {Element} header
     */
    _handleHeaderClick: function(header) {
      const accordion = this._getParentAccordion(header);

      if (!accordion) {
        console.warn('Accordion: No parent accordion found');
        return;
      }

      const isExpanded = accordion.classList.contains(this.EXPANDED_CLASS);

      // Fecha todos os accordions
      this._closeAll();

      // Abre o clicado se não estava expandido
      if (!isExpanded) {
        this._openAccordion(accordion, header);
      }

      // Foca no header para acessibilidade
      if (typeof header.focus === 'function') {
        header.focus();
      }
    },

    /**
     * Obtém accordion pai
     * @private
     * @param {Element} element
     * @returns {Element|null}
     */
    _getParentAccordion: function(element) {
      let current = element;

      while (current && current !== document.body) {
        if (current.classList && current.classList.contains('accordion')) {
          return current;
        }
        current = current.parentElement;
      }

      return null;
    },

    /**
     * Fecha todos os accordions
     * @private
     */
    _closeAll: function() {
      if (!this._elements.accordions) return;

      for (let i = 0; i < this._elements.accordions.length; i++) {
        const accordion = this._elements.accordions[i];
        const header = accordion.querySelector(this.HEADER_SELECTOR);

        accordion.classList.remove(this.EXPANDED_CLASS);

        if (header) {
          header.setAttribute(this.ARIA_EXPANDED, 'false');
        }
      }
    },

    /**
     * Abre um accordion específico
     * @private
     * @param {Element} accordion
     * @param {Element} header
     */
    _openAccordion: function(accordion, header) {
      accordion.classList.add(this.EXPANDED_CLASS);
      header.setAttribute(this.ARIA_EXPANDED, 'true');
    },

    /**
     * Abre accordion por ID
     * @param {string} id
     * @returns {boolean}
     */
    openById: function(id) {
      if (!App.Utils.isValidString(id)) return false;

      const accordion = document.getElementById(id);

      if (!accordion || !accordion.classList.contains('accordion')) {
        return false;
      }

      const header = accordion.querySelector(this.HEADER_SELECTOR);

      if (!header) return false;

      this._closeAll();
      this._openAccordion(accordion, header);

      return true;
    },

    /**
     * Fecha accordion por ID
     * @param {string} id
     * @returns {boolean}
     */
    closeById: function(id) {
      if (!App.Utils.isValidString(id)) return false;

      const accordion = document.getElementById(id);

      if (!accordion || !accordion.classList.contains('accordion')) {
        return false;
      }

      const header = accordion.querySelector(this.HEADER_SELECTOR);

      if (header) {
        header.setAttribute(this.ARIA_EXPANDED, 'false');
      }

      accordion.classList.remove(this.EXPANDED_CLASS);

      return true;
    },

    /**
     * Recarrega cache (útil após DOM dinâmico)
     */
    refresh: function() {
      this._cacheElements();
    }
  };

})();
