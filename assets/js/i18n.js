/**
 * Internacionalização (i18n) - PT/EN
 * @module I18n
 * @requires App.Utils
 */

(function() {
  'use strict';

  window.App = window.App || {};

  /**
   * Gerenciador de internacionalização
   * @namespace App.I18n
   */
  App.I18n = {
    /**
     * Chave do localStorage
     * @constant {string}
     */
    STORAGE_KEY: 'lang',

    /**
     * Valores válidos de idioma
     * @constant {Array<string>}
     */
    VALID_LANGS: ['pt', 'en'],

    /**
     * Idioma padrão
     * @constant {string}
     */
    DEFAULT_LANG: 'pt',

    /**
     * Mapeamento de códigos de idioma para atributo lang
     * @constant {Object}
     */
    LANG_ATTR_MAP: {
      'pt': 'pt-BR',
      'en': 'en'
    },

    /**
     * Elementos DOM cacheados
     * @private
     */
    _elements: {
      html: null,
      toggle: null,
      text: null
    },

    /**
     * Cache de elementos de idioma
     * @private
     */
    _langElements: {
      pt: null,
      en: null
    },

    /**
     * Inicializa o módulo i18n
     * @returns {boolean}
     */
    init: function() {
      try {
        this._cacheElements();

        if (!this._validateElements()) {
          console.warn('I18n: Required elements not found');
          return false;
        }

        this._cacheLangElements();
        this._bindEvents();
        this._loadSavedLanguage();

        return true;
      } catch (error) {
        console.error('I18n initialization failed:', error);
        return false;
      }
    },

    /**
     * Cache de elementos principais
     * @private
     */
    _cacheElements: function() {
      this._elements.html = document.documentElement;
      this._elements.toggle = document.getElementById('lang-toggle');
      this._elements.text = document.getElementById('lang-text');
    },

    /**
     * Cache de elementos de idioma
     * @private
     */
    _cacheLangElements: function() {
      this._langElements.pt = document.querySelectorAll('.lang-pt');
      this._langElements.en = document.querySelectorAll('.lang-en');
    },

    /**
     * Validação de elementos
     * @private
     * @returns {boolean}
     */
    _validateElements: function() {
      return App.Utils.isValidElement(this._elements.html) &&
             App.Utils.isValidElement(this._elements.toggle) &&
             App.Utils.isValidElement(this._elements.text);
    },

    /**
     * Vincula eventos
     * @private
     */
    _bindEvents: function() {
      const self = this;

      this._elements.toggle.addEventListener('click', function(event) {
        event.preventDefault();
        self._handleToggle();
      });

      this._elements.toggle.addEventListener('keydown', function(event) {
        if (event.key === 'Enter' || event.key === ' ') {
          event.preventDefault();
          self._handleToggle();
        }
      });
    },

    /**
     * Carrega idioma salvo
     * @private
     */
    _loadSavedLanguage: function() {
      const savedLang = App.Utils.Storage.get(this.STORAGE_KEY, this.DEFAULT_LANG);
      const validLang = this._validateLang(savedLang);

      this._applyLanguage(validLang);
    },

    /**
     * Valida código de idioma
     * @private
     * @param {*} lang
     * @returns {string}
     */
    _validateLang: function(lang) {
      if (!App.Utils.isValidString(lang)) {
        return this.DEFAULT_LANG;
      }

      const normalized = lang.toLowerCase().trim();

      if (this.VALID_LANGS.indexOf(normalized) === -1) {
        return this.DEFAULT_LANG;
      }

      return normalized;
    },

    /**
     * Handler do toggle
     * @private
     */
    _handleToggle: function() {
      const currentLang = this.getCurrentLang();
      const newLang = currentLang === 'pt' ? 'en' : 'pt';

      this._applyLanguage(newLang);
      this._saveLanguage(newLang);
    },

    /**
     * Aplica idioma ao DOM
     * @private
     * @param {string} lang
     */
    _applyLanguage: function(lang) {
      const safeLang = this._validateLang(lang);
      const attrLang = this.LANG_ATTR_MAP[safeLang] || safeLang;

      // Atualiza atributo lang do HTML
      this._elements.html.setAttribute('lang', attrLang);

      // Atualiza texto do botão
      this._updateButtonText(safeLang);

      // Alterna visibilidade dos elementos
      this._toggleElements(safeLang);
    },

    /**
     * Atualiza texto do botão
     * @private
     * @param {string} lang
     */
    _updateButtonText: function(lang) {
      // Mostra o idioma oposto ao atual
      const buttonText = lang === 'pt' ? 'EN' : 'PT';
      this._elements.text.textContent = buttonText;

      // Atualiza aria-label para acessibilidade
      const ariaLabel = lang === 'pt' ? 'Mudar idioma para inglês' : 'Change language to Portuguese';
      this._elements.toggle.setAttribute('aria-label', ariaLabel);
    },

    /**
     * Alterna visibilidade dos elementos de idioma
     * @private
     * @param {string} activeLang
     */
    _toggleElements: function(activeLang) {
      const inactiveLang = activeLang === 'pt' ? 'en' : 'pt';

      // Mostra elementos do idioma ativo
      if (this._langElements[activeLang]) {
        for (let i = 0; i < this._langElements[activeLang].length; i++) {
          this._langElements[activeLang][i].style.display = '';
        }
      }

      // Esconde elementos do idioma inativo
      if (this._langElements[inactiveLang]) {
        for (let i = 0; i < this._langElements[inactiveLang].length; i++) {
          this._langElements[inactiveLang][i].style.display = 'none';
        }
      }
    },

    /**
     * Salva idioma no storage
     * @private
     * @param {string} lang
     */
    _saveLanguage: function(lang) {
      App.Utils.Storage.set(this.STORAGE_KEY, lang);
    },

    /**
     * Define idioma programaticamente
     * @param {string} lang
     */
    setLanguage: function(lang) {
      const safeLang = this._validateLang(lang);
      this._applyLanguage(safeLang);
      this._saveLanguage(safeLang);
    },

    /**
     * Retorna idioma atual
     * @returns {string}
     */
    getCurrentLang: function() {
      const current = this._elements.html.getAttribute('lang');
      // Converte pt-BR para pt
      if (current === 'pt-BR') return 'pt';
      return this._validateLang(current);
    },

    /**
     * Recarrega cache de elementos (útil após DOM dinâmico)
     */
    refresh: function() {
      this._cacheLangElements();
      this._applyLanguage(this.getCurrentLang());
    }
  };

})();
