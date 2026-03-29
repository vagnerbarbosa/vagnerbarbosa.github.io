/**
 * Gerenciamento de Tema (Light/Dark)
 * @module Theme
 * @requires App.Utils
 */

(function() {
  'use strict';

  window.App = window.App || {};

  /**
   * Gerenciador de tema
   * @namespace App.Theme
   */
  App.Theme = {
    /**
     * Chave do localStorage
     * @constant {string}
     */
    STORAGE_KEY: 'theme',

    /**
     * Valores válidos de tema
     * @constant {Array<string>}
     */
    VALID_THEMES: ['light', 'dark'],

    /**
     * Valor padrão
     * @constant {string}
     */
    DEFAULT_THEME: 'light',

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
     * Inicializa o módulo de tema
     * @returns {boolean} Sucesso da inicialização
     */
    init: function() {
      try {
        this._cacheElements();

        if (!this._validateElements()) {
          console.warn('Theme: Required elements not found');
          return false;
        }

        this._bindEvents();
        this._loadSavedTheme();

        return true;
      } catch (error) {
        console.error('Theme initialization failed:', error);
        return false;
      }
    },

    /**
     * Cache de elementos DOM
     * @private
     */
    _cacheElements: function() {
      this._elements.html = document.documentElement;
      this._elements.toggle = document.getElementById('theme-toggle');
      this._elements.textPt = document.getElementById('theme-text-pt');
      this._elements.textEn = document.getElementById('theme-text-en');
    },

    /**
     * Validação de elementos
     * @private
     * @returns {boolean}
     */
    _validateElements: function() {
      return App.Utils.isValidElement(this._elements.html) &&
             App.Utils.isValidElement(this._elements.toggle);
    },

    /**
     * Vincula eventos
     * @private
     */
    _bindEvents: function() {
      const self = this;

      // Usa event delegation para melhor performance
      this._elements.toggle.addEventListener('click', function(event) {
        event.preventDefault();
        self._handleToggle();
      });

      // Previne execução múltipla via teclado
      this._elements.toggle.addEventListener('keydown', function(event) {
        if (event.key === 'Enter' || event.key === ' ') {
          event.preventDefault();
          self._handleToggle();
        }
      });

      // Atualiza quando o idioma muda
      document.addEventListener('languageChanged', function() {
        self._updateText(self.getCurrentTheme());
      });
    },

    /**
     * Carrega tema salvo
     * @private
     */
    _loadSavedTheme: function() {
      const savedTheme = App.Utils.Storage.get(this.STORAGE_KEY, this.DEFAULT_THEME);
      const validTheme = this._validateTheme(savedTheme);

      this._applyTheme(validTheme);
    },

    /**
     * Valida valor do tema
     * @private
     * @param {string} theme
     * @returns {string} Tema validado
     */
    _validateTheme: function(theme) {
      if (!App.Utils.isValidString(theme)) {
        return this.DEFAULT_THEME;
      }

      const normalized = theme.toLowerCase().trim();

      if (this.VALID_THEMES.indexOf(normalized) === -1) {
        return this.DEFAULT_THEME;
      }

      return normalized;
    },

    /**
     * Handler do toggle
     * @private
     */
    _handleToggle: function() {
      const currentTheme = this._elements.html.getAttribute('data-theme');
      const current = this._validateTheme(currentTheme);
      const newTheme = current === 'light' ? 'dark' : 'light';

      this._applyTheme(newTheme);
      this._saveTheme(newTheme);
    },

    /**
     * Aplica tema ao DOM
     * @private
     * @param {string} theme
     */
    _applyTheme: function(theme) {
      // Sanitiza antes de aplicar
      const safeTheme = this._validateTheme(theme);

      this._elements.html.setAttribute('data-theme', safeTheme);
      this._updateText(safeTheme);
    },

    /**
     * Atualiza texto do botão
     * @private
     * @param {string} theme
     */
    _updateText: function(theme) {
      // Verifica idioma atual para textos bilíngues
      // Só usa I18n se estiver totalmente inicializado
      let currentLang = 'pt';
      if (window.App.I18n && window.App.I18n._elements && window.App.I18n._elements.html) {
        currentLang = window.App.I18n.getCurrentLang();
      }
      const isEnglish = currentLang === 'en';

      // Textos bilíngues para o botão de tema
      const textPt = theme === 'light' ? 'Modo escuro' : 'Modo claro';
      const textEn = theme === 'light' ? 'Dark mode' : 'Light mode';

      // Atualiza ambos os spans se existirem
      if (App.Utils.isValidElement(this._elements.textPt)) {
        this._elements.textPt.textContent = textPt;
      }
      if (App.Utils.isValidElement(this._elements.textEn)) {
        this._elements.textEn.textContent = textEn;
      }

      // Atualiza aria-pressed e aria-label bilíngue para acessibilidade
      this._elements.toggle.setAttribute('aria-pressed', theme === 'dark');
      const ariaLabel = isEnglish
        ? (theme === 'light' ? 'Enable dark mode' : 'Enable light mode')
        : (theme === 'light' ? 'Ativar modo escuro' : 'Ativar modo claro');
      this._elements.toggle.setAttribute('aria-label', ariaLabel);
    },

    /**
     * Salva tema no storage
     * @private
     * @param {string} theme
     */
    _saveTheme: function(theme) {
      App.Utils.Storage.set(this.STORAGE_KEY, theme);
    },

    /**
     * Define tema programaticamente
     * @param {string} theme
     */
    setTheme: function(theme) {
      const safeTheme = this._validateTheme(theme);
      this._applyTheme(safeTheme);
      this._saveTheme(safeTheme);
    },

    /**
     * Retorna tema atual
     * @returns {string}
     */
    getCurrentTheme: function() {
      const current = this._elements.html.getAttribute('data-theme');
      return this._validateTheme(current);
    }
  };

})();
