/**
 * Fade In Animation Module
 * Animates elements with fade-in effect when they become visible
 * Matches annamona.co implementation
 * @version 2.0.0
 */

(function() {
  'use strict';

  /**
   * FadeIn module
   * @namespace App.FadeIn
   */
  window.App = window.App || {};

  const FadeIn = {
    /**
     * Initialize the fade-in animations
     * Handles prefers-reduced-motion for accessibility
     */
    init: function() {
      // Check if user prefers reduced motion
      if (window.matchMedia('(prefers-reduced-motion: reduce)').matches) {
        this.showAll();
      }
      // Otherwise, CSS animation runs automatically via @keyframes
    },

    /**
     * Show all elements immediately (fallback for reduced motion)
     */
    showAll: function() {
      const elements = document.querySelectorAll('.fade-in');
      elements.forEach(function(el) {
        el.style.opacity = '1';
        el.style.animation = 'none';
      });
    }
  };

  // Expose module
  window.App.FadeIn = FadeIn;

  // Auto-initialize when DOM is ready
  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', function() {
      FadeIn.init();
    });
  } else {
    FadeIn.init();
  }

})();
