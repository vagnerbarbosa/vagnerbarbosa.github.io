/**
 * Utilitários seguros para o site
 * @module Utils
 * @version 1.0.0
 */

(function() {
  'use strict';

  /**
   * Namespace global do aplicativo
   * @namespace
   */
  window.App = window.App || {};

  /**
   * Utilitários de sanitização e segurança
   * @namespace App.Utils
   */
  App.Utils = {
    /**
     * Inicializa o módulo de utilitários
     * @returns {boolean} Sempre retorna true
     */
    init: function() {
      // Utils é um módulo utilitário, não requer inicialização
      return true;
    },

    /**
     * Sanitiza uma string para prevenir XSS
     * @param {string} str - String a ser sanitizada
     * @returns {string} String sanitizada
     */
    escapeHtml: function(str) {
      if (typeof str !== 'string') return '';
      const div = document.createElement('div');
      div.textContent = str;
      return div.innerHTML;
    },

    /**
     * Sanitiza atributos HTML
     * @param {string} str - String a ser sanitizada
     * @returns {string} String sanitizada para atributos
     */
    escapeAttr: function(str) {
      if (typeof str !== 'string') return '';
      return str
        .replace(/&/g, '&amp;')
        .replace(/"/g, '&quot;')
        .replace(/'/g, '&#39;')
        .replace(/</g, '&lt;')
        .replace(/>/g, '&gt;');
    },

    /**
     * Verifica se uma string é válida e segura
     * @param {*} str - Valor a ser verificado
     * @returns {boolean}
     */
    isValidString: function(str) {
      return typeof str === 'string' && str.trim().length > 0;
    },

    /**
     * Verifica se um elemento DOM existe e é válido
     * @param {Element|null} element - Elemento a ser verificado
     * @returns {boolean}
     */
    isValidElement: function(element) {
      return element instanceof Element;
    },

    /**
     * Safe LocalStorage wrapper com tratamento de erros
     * @namespace App.Utils.Storage
     */
    Storage: {
      /**
       * Verifica se localStorage está disponível
       * @returns {boolean}
       */
      isAvailable: function() {
        try {
          const test = '__storage_test__';
          localStorage.setItem(test, test);
          localStorage.removeItem(test);
          return true;
        } catch (e) {
          return false;
        }
      },

      /**
       * Salva um valor no localStorage de forma segura
       * @param {string} key - Chave
       * @param {string} value - Valor
       * @returns {boolean} Sucesso da operação
       */
      set: function(key, value) {
        try {
          if (!this.isAvailable()) return false;
          if (!App.Utils.isValidString(key)) return false;

          // Validação de tamanho para prevenir abuse
          if (key.length > 100) return false;
          if (value.length > 10000) return false;

          localStorage.setItem(key, value);
          return true;
        } catch (e) {
          console.warn('Storage set failed:', e);
          return false;
        }
      },

      /**
       * Recupera um valor do localStorage
       * @param {string} key - Chave
       * @param {string} [defaultValue=''] - Valor padrão
       * @returns {string} Valor recuperado ou padrão
       */
      get: function(key, defaultValue) {
        defaultValue = defaultValue || '';
        try {
          if (!this.isAvailable()) return defaultValue;
          if (!App.Utils.isValidString(key)) return defaultValue;

          const value = localStorage.getItem(key);
          return value !== null ? value : defaultValue;
        } catch (e) {
          console.warn('Storage get failed:', e);
          return defaultValue;
        }
      },

      /**
       * Remove um item do localStorage
       * @param {string} key - Chave
       * @returns {boolean}
       */
      remove: function(key) {
        try {
          if (!this.isAvailable()) return false;
          localStorage.removeItem(key);
          return true;
        } catch (e) {
          console.warn('Storage remove failed:', e);
          return false;
        }
      }
    },

    /**
     * Debounce para eventos frequentes
     * @param {Function} func - Função a ser executada
     * @param {number} wait - Tempo de espera em ms
     * @returns {Function}
     */
    debounce: function(func, wait) {
      let timeout;
      return function executedFunction() {
        const context = this;
        const args = arguments;
        const later = function() {
          timeout = null;
          func.apply(context, args);
        };
        clearTimeout(timeout);
        timeout = setTimeout(later, wait);
      };
    },

    /**
     * Executa callback apenas quando DOM está pronto
     * @param {Function} callback - Função a ser executada
     */
    onReady: function(callback) {
      if (document.readyState === 'loading') {
        document.addEventListener('DOMContentLoaded', callback);
      } else {
        callback();
      }
    }
  };

})();
