<template>
  <div class="app">
    <!-- 1. Навбар -->
    <header class="navbar">
      <div class="container navbar-content">
        <div class="logo">
          <h1>QAZAQ PIPE</h1>
          <span>ПОЛИЭТИЛЕНОВЫЕ ТРУБЫ</span>
        </div>

        <nav class="nav-links" :class="{ 'nav-active': isMenuOpen }">
          <a href="#hero" @click="closeMenu">{{ t.home }}</a>
          <a href="#about" @click="closeMenu">{{ t.about }}</a>
          <a href="#production" @click="closeMenu">{{ t.production }}</a>
          <a href="#catalog" @click="closeMenu">{{ t.catalog }}</a>
          <a href="#calculator" @click="closeMenu">{{ t.calc }}</a>
          <a href="#certificates" @click="closeMenu">{{ t.certificates }}</a>
          <a href="#advantages" @click="closeMenu">{{ t.advantages }}</a>
          <a href="#contacts" @click="closeMenu">{{ t.contacts }}</a>
          <a :href="getWhatsappUrl('Запрос КП')" target="_blank" class="btn-kp mobile-kp" @click="closeMenu">{{ t.btnKp }}</a>
        </nav>

        <div class="nav-right">
          <a :href="getWhatsappUrl('Запрос КП')" target="_blank" class="btn-kp desktop-kp">{{ t.btnKp }}</a>
          <!-- Кнопка переключения языка -->
          <button class="lang-switch" @click="toggleLanguage" type="button" :title="currentLang === 'RU' ? 'Қазақ тіліне ауыстыру' : 'Переключить на русский'">
            {{ currentLang === 'RU' ? 'KZ' : 'RU' }}
          </button>
          <!-- Кнопка бургер-меню справа -->
          <button class="burger-btn" @click="isMenuOpen = !isMenuOpen" aria-label="Меню">
            <span :class="{ 'open': isMenuOpen }"></span>
            <span :class="{ 'open': isMenuOpen }"></span>
            <span :class="{ 'open': isMenuOpen }"></span>
          </button>
        </div>
      </div>
    </header>

    <!-- 2. Hero Баннер -->
    <section id="hero" class="hero">
      <div class="hero-overlay"></div>
      <div class="container hero-container">
        <div class="hero-text-block">
          <h2>{{ t.heroTitle }}</h2>
          <p class="subtitle">{{ t.heroSubtitle }}</p>
          <a :href="getWhatsappUrl('Здравствуйте! Хочу получить коммерческое предложение.')" target="_blank" class="btn-hero">
            {{ t.btnHero }}
          </a>
        </div>

        <!-- 4 Плашки показателей -->
        <div class="hero-stats-bar">
          <div class="stat-box">
            <div class="stat-icon">
              <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect><line x1="16" y1="2" x2="16" y2="6"></line><line x1="8" y1="2" x2="8" y2="6"></line><line x1="3" y1="10" x2="21" y2="10"></line></svg>
            </div>
            <div>
              <strong>{{ t.stat1Title }}</strong>
              <span>{{ t.stat1Desc }}</span>
            </div>
          </div>

          <div class="stat-box">
            <div class="stat-icon">
              <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M2 20a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V8l-7 5V8l-7 5V4a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2Z"></path><path d="M17 18h1"></path><path d="M12 18h1"></path><path d="M7 18h1"></path></svg>
            </div>
            <div>
              <strong>{{ t.stat2Title }}</strong>
              <span>{{ t.stat2Desc }}</span>
            </div>
          </div>

          <div class="stat-box">
            <div class="stat-icon">
              <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"></path><polyline points="3.27 6.96 12 12.01 20.73 6.96"></polyline><line x1="12" y1="22.08" x2="12" y2="12"></line></svg>
            </div>
            <div>
              <strong>{{ t.stat3Title }}</strong>
              <span>{{ t.stat3Desc }}</span>
            </div>
          </div>

          <div class="stat-box">
            <div class="stat-icon">
              <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M10 2v7.31"></path><path d="M14 9.31V2"></path><path d="M8.5 2h7"></path><path d="M14 12.5a3 3 0 1 1-4 0l-3.3-6.6a1 1 0 0 1 .9-1.4h9.8a1 1 0 0 1 .9 1.4Z"></path></svg>
            </div>
            <div>
              <strong>{{ t.stat4Title }}</strong>
              <span>{{ t.stat4Desc }}</span>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- 3. Блок: О Компании -->
    <section id="about" class="section-about">
      <div class="container grid-about">
        <div class="about-text">
          <span class="sub-label">{{ t.aboutSub }}</span>
          <h3>{{ t.aboutTitle }}</h3>
          <p>{{ t.aboutDesc1 }}</p>
          <p>{{ t.aboutDesc2 }}</p>
          <a href="#production" class="btn-about">{{ t.aboutBtn }}</a>
        </div>
        <div class="about-img-box">
          <img src="/about_company.png" alt="Здание завода Qazaq Pipe" class="about-img">
        </div>
      </div>
    </section>

    <!-- 4. Блок: Технологическая линия производства -->
    <section id="production" class="section-white">
      <div class="container">
        <h3 class="section-title center">{{ t.techTitle }}</h3>
        <p class="section-subtitle center">{{ t.techSubtitle }}</p>

        <div class="production-carousel">
          <button
            class="carousel-arrow carousel-arrow-left"
            type="button"
            aria-label="Предыдущий этап"
            @click="prevSlide"
          >
            ‹
          </button>

          <div ref="carouselWindow" class="carousel-window">
            <div
              class="carousel-track"
              :class="{ 'no-transition': isResettingCarousel }"
              :style="carouselTrackStyle"
            >
              <div
                v-for="(step, index) in carouselSlides"
                :key="`${step.title}-${index}`"
                class="step-item production-slide"
              >
                <div class="step-num">{{ step.number }}</div>

                <div class="step-img-box production-image-box">
                  <img :src="step.image" :alt="step.title" class="production-image">
                </div>

                <h5>{{ step.title }}</h5>
                <p>{{ step.description }}</p>
              </div>
            </div>
          </div>

          <button
            class="carousel-arrow carousel-arrow-right"
            type="button"
            aria-label="Следующий этап"
            @click="nextSlide"
          >
            ›
          </button>
        </div>

        <div class="carousel-dots">
          <button
            v-for="(_, index) in productionSteps"
            :key="index"
            type="button"
            class="carousel-dot"
            :class="{ active: activeDot === index }"
            :aria-label="`Перейти к этапу ${index + 1}`"
            @click="goToSlide(index)"
          ></button>
        </div>
      </div>
    </section>

    <!-- 5. Блок: Наша Продукция -->
    <section id="catalog" class="section-gray">
      <div class="container">
        <h3 class="section-title center">{{ t.prodTitle }}</h3>
        <div class="product-cards-grid">
          <div class="prod-card animated-card">
            <div class="prod-img-container">
              <img src="/Water_supply_pipe.png" :alt="t.p1Title" class="prod-img">
            </div>
            <h4>{{ t.p1Title }}</h4>
            <p>{{ t.p1Desc }}</p>
          </div>
          <div class="prod-card animated-card">
            <div class="prod-img-container">
              <img src="/Gas_supply_pipes.png" :alt="t.p2Title" class="prod-img">
            </div>
            <h4>{{ t.p2Title }}</h4>
            <p>{{ t.p2Desc }}</p>
          </div>
          <div class="prod-card animated-card">
            <div class="prod-img-container">
              <img src="/Sewerage_pipes.png" :alt="t.p3Title" class="prod-img">
            </div>
            <h4>{{ t.p3Title }}</h4>
            <p>{{ t.p3Desc }}</p>
          </div>
          <div class="prod-card animated-card">
            <div class="prod-img-container">
              <img src="/Industrial_Piping_Systems.png" :alt="t.p4Title" class="prod-img">
            </div>
            <h4>{{ t.p4Title }}</h4>
            <p>{{ t.p4Desc }}</p>
          </div>
          <div class="prod-card animated-card">
            <div class="prod-img-container">
              <img src="/Irrigation_pipes.png" :alt="t.p5Title" class="prod-img">
            </div>
            <h4>{{ t.p5Title }}</h4>
            <p>{{ t.p5Desc }}</p>
          </div>
          <div class="prod-card animated-card">
            <div class="prod-img-container">
              <img src="/Casing_pipes.png" :alt="t.p6Title" class="prod-img">
            </div>
            <h4>{{ t.p6Title }}</h4>
            <p>{{ t.p6Desc }}</p>
          </div>
        </div>

        <!-- Справочная таблица параметров -->
        <div class="table-container">
          <h4>{{ t.tableTitle }}</h4>
          <div class="table-wrapper">
            <table class="spec-table">
              <thead>
                <tr>
                  <th>{{ t.tableTh1 }}</th>
                  <th>63</th>
                  <th>90</th>
                  <th>110</th>
                  <th>160</th>
                  <th>225</th>
                  <th>315</th>
                  <th>400</th>
                  <th>500</th>
                </tr>
              </thead>
              <tbody>
                <tr><td><strong>SDR 11</strong></td><td>5.8</td><td>8.2</td><td>10.0</td><td>14.6</td><td>20.5</td><td>28.6</td><td>36.3</td><td>45.4</td></tr>
                <tr><td><strong>SDR 17</strong></td><td>3.8</td><td>5.4</td><td>6.6</td><td>9.5</td><td>13.4</td><td>18.7</td><td>23.7</td><td>29.7</td></tr>
                <tr><td><strong>SDR 21</strong></td><td>3.0</td><td>4.3</td><td>5.3</td><td>7.7</td><td>10.8</td><td>15.1</td><td>19.1</td><td>23.9</td></tr>
                <tr><td><strong>SDR 26</strong></td><td>2.4</td><td>3.5</td><td>4.2</td><td>6.2</td><td>8.6</td><td>12.1</td><td>15.3</td><td>19.1</td></tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </section>

    <!-- 5. Блок: Калькулятор труб -->
    <section id="calculator" class="section-calculator">
      <div class="container">
        <h3 class="section-title center">{{ t.calcTitle }}</h3>
        <p class="section-subtitle center">{{ t.calcSubtitle }}</p>

        <div class="calc-card">
          <div class="calc-grid">
            <div class="calc-field">
              <label>{{ t.calcDiamLabel }}</label>
              <select v-model.number="calcDiameter">
                <option v-for="d in diameterOptions" :key="d" :value="d">{{ d }} мм</option>
              </select>
            </div>

            <div class="calc-field">
              <label>{{ t.calcSdrLabel }}</label>
              <select v-model.number="calcSdr">
                <option v-for="sdr in sdrOptions" :key="sdr" :value="sdr">SDR {{ sdr }}</option>
              </select>
            </div>

            <div class="calc-field">
              <label>{{ t.calcMetersLabel }}</label>
              <input type="number" v-model.number="calcMeters" min="1" :placeholder="t.calcMetersPlaceholder">
            </div>
          </div>

          <div class="calc-result">
            <div class="res-item">
              <span>{{ t.calcRes1 }}</span>
              <strong>{{ weightPerMeter }} кг/м</strong>
            </div>
            <div class="res-item main-res">
              <span>{{ t.calcRes2 }}</span>
              <strong>{{ totalWeight }} кг ({{ (totalWeight / 1000).toFixed(2) }} т)</strong>
            </div>
            <a :href="getWhatsappUrl(`Расчет: Диаметр ${calcDiameter}мм, SDR ${calcSdr}, Длина ${calcMeters}м. Общий вес около ${totalWeight} кг.`)" target="_blank" class="btn-calc">
              {{ t.calcBtn }}
            </a>
          </div>
        </div>
      </div>
    </section>

    <!-- 7. Блок: Сертификаты -->
    <section id="certificates" class="section-certificates">
      <div class="container">
        <h3 class="section-title center">{{ t.certTitle }}</h3>
        <p class="section-subtitle center">{{ t.certSubtitle }}</p>
        <div class="cert-grid">
          <a
            v-for="(cert, index) in certificates"
            :key="index"
            :href="cert.file"
            target="_blank"
            rel="noopener noreferrer"
            class="cert-card animated-card"
          >
            <div class="cert-icon">
              <svg width="40" height="40" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg">
                <rect x="10" y="6" width="28" height="36" rx="3" fill="#0284c7" fill-opacity="0.12" stroke="#0284c7" stroke-width="2"/>
                <path d="M18 16h12M18 22h12M18 28h8" stroke="#0284c7" stroke-width="2" stroke-linecap="round"/>
                <circle cx="34" cy="36" r="8" fill="#0284c7"/>
                <path d="M31 36l2 2 4-4" stroke="white" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
              </svg>
            </div>
            <div class="cert-body">
              <span class="cert-lang" :class="cert.lang === 'KZ' ? 'lang-kz' : 'lang-ru'">{{ cert.lang }}</span>
              <h5>{{ cert.title }}</h5>
              <p>{{ cert.desc }}</p>
            </div>
            <span class="cert-download">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/>
                <polyline points="7 10 12 15 17 10"/>
                <line x1="12" y1="15" x2="12" y2="3"/>
              </svg>
              PDF
            </span>
          </a>
        </div>
      </div>
    </section>

    <!-- 8. Блок: Почему выбирают QAZAQ PIPE -->
    <section id="advantages" class="section-dark">
      <div class="container">
        <h3 class="section-title center white-text">{{ t.advTitle }}</h3>
        <div class="adv-grid">
          <!-- 1. Собственное производство -->
          <div class="adv-item">
            <div class="adv-icon">
              <svg width="48" height="48" viewBox="0 0 64 64" fill="none" xmlns="http://www.w3.org/2000/svg">
                <!-- Основное здание -->
                <rect x="8" y="28" width="28" height="28" rx="2" fill="#F472B6"/>
                <!-- Правое здание -->
                <rect x="36" y="20" width="20" height="36" rx="2" fill="#EC4899"/>
                <!-- Трубы/дымоходы -->
                <rect x="14" y="16" width="6" height="12" rx="1" fill="#F9A8D4"/>
                <rect x="24" y="12" width="6" height="16" rx="1" fill="#F9A8D4"/>
                <rect x="42" y="10" width="5" height="10" rx="1" fill="#FBCFE8"/>
                <!-- Окна -->
                <rect x="13" y="34" width="6" height="6" rx="1" fill="#0b1329"/>
                <rect x="23" y="34" width="6" height="6" rx="1" fill="#0b1329"/>
                <rect x="13" y="44" width="6" height="6" rx="1" fill="#0b1329"/>
                <rect x="23" y="44" width="6" height="6" rx="1" fill="#0b1329"/>
                <rect x="41" y="28" width="5" height="5" rx="1" fill="#0b1329"/>
                <rect x="49" y="28" width="5" height="5" rx="1" fill="#0b1329"/>
                <rect x="41" y="38" width="5" height="5" rx="1" fill="#0b1329"/>
                <rect x="49" y="38" width="5" height="5" rx="1" fill="#0b1329"/>
                <!-- Дверь -->
                <rect x="44" y="46" width="8" height="10" rx="1" fill="#0b1329"/>
              </svg>
            </div>
            <h5>{{ t.adv1Title }}</h5>
            <p>{{ t.adv1Desc }}</p>
          </div>

          <!-- 2. Немецкое оборудование -->
          <div class="adv-item">
            <div class="adv-icon">
              <svg width="48" height="48" viewBox="0 0 64 64" fill="none" xmlns="http://www.w3.org/2000/svg">
                <!-- Шестерёнка -->
                <path d="M32 10l3.2 6.2 6.9-1.1 1.1 6.9 6.2 3.2-3.2 6.2 4.8 5.1-4.8 5.1 3.2 6.2-6.2 3.2-1.1 6.9-6.9-1.1L32 54l-3.2-6.2-6.9 1.1-1.1-6.9-6.2-3.2 3.2-6.2-4.8-5.1 4.8-5.1-3.2-6.2 6.2-3.2 1.1-6.9 6.9 1.1L32 10z" fill="#34D399"/>
                <circle cx="32" cy="32" r="10" fill="#0b1329"/>
                <circle cx="32" cy="32" r="5.5" fill="#6EE7B7"/>
              </svg>
            </div>
            <h5>{{ t.adv2Title }}</h5>
            <p>{{ t.adv2Desc }}</p>
          </div>

          <!-- 3. Стабильное качество -->
          <div class="adv-item">
            <div class="adv-icon">
              <svg width="48" height="48" viewBox="0 0 64 64" fill="none" xmlns="http://www.w3.org/2000/svg">
                <!-- Щит -->
                <path d="M32 6C32 6 12 14 12 14v16c0 14.5 9.5 24.5 20 28 10.5-3.5 20-13.5 20-28V14S32 6 32 6z" fill="#60A5FA"/>
                <!-- Галочка -->
                <path d="M22 32l7 7 13-14" stroke="#0b1329" stroke-width="4.5" stroke-linecap="round" stroke-linejoin="round"/>
              </svg>
            </div>
            <h5>{{ t.adv3Title }}</h5>
            <p>{{ t.adv3Desc }}</p>
          </div>

          <!-- 4. До 7000 тонн -->
          <div class="adv-item">
            <div class="adv-icon">
              <svg width="48" height="48" viewBox="0 0 64 64" fill="none" xmlns="http://www.w3.org/2000/svg">
                <!-- Столбцы диаграммы -->
                <rect x="10" y="34" width="12" height="20" rx="2.5" fill="#FBBF24"/>
                <rect x="26" y="22" width="12" height="32" rx="2.5" fill="#34D399"/>
                <rect x="42" y="12" width="12" height="42" rx="2.5" fill="#60A5FA"/>
                <!-- Блики -->
                <rect x="12" y="36" width="4" height="8" rx="1" fill="#FDE68A" fill-opacity="0.6"/>
                <rect x="28" y="24" width="4" height="10" rx="1" fill="#6EE7B7" fill-opacity="0.6"/>
                <rect x="44" y="14" width="4" height="12" rx="1" fill="#93C5FD" fill-opacity="0.6"/>
              </svg>
            </div>
            <h5>{{ t.adv4Title }}</h5>
            <p>{{ t.adv4Desc }}</p>
          </div>

          <!-- 5. Диаметры -->
          <div class="adv-item">
            <div class="adv-icon">
              <svg width="48" height="48" viewBox="0 0 64 64" fill="none" xmlns="http://www.w3.org/2000/svg">
                <!-- Линейка (диагональная) -->
                <rect x="10" y="28" width="44" height="12" rx="2.5" transform="rotate(-20 32 34)" fill="#38BDF8"/>
                <!-- Деления -->
                <g transform="rotate(-20 32 34)">
                  <rect x="16" y="28" width="2" height="5" fill="#0b1329"/>
                  <rect x="22" y="28" width="2" height="7" fill="#0b1329"/>
                  <rect x="28" y="28" width="2" height="5" fill="#0b1329"/>
                  <rect x="34" y="28" width="2" height="7" fill="#0b1329"/>
                  <rect x="40" y="28" width="2" height="5" fill="#0b1329"/>
                  <rect x="46" y="28" width="2" height="7" fill="#0b1329"/>
                </g>
              </svg>
            </div>
            <h5>{{ t.adv5Title }}</h5>
            <p>{{ t.adv5Desc }}</p>
          </div>

          <!-- 6. Поставка -->
          <div class="adv-item">
            <div class="adv-icon">
              <svg width="48" height="48" viewBox="0 0 64 64" fill="none" xmlns="http://www.w3.org/2000/svg">
                <!-- Кузов -->
                <rect x="6" y="22" width="32" height="22" rx="2.5" fill="#FB923C"/>
                <!-- Кабина -->
                <path d="M38 28h12l6 10v10H38V28z" fill="#F97316"/>
                <!-- Окно кабины -->
                <rect x="42" y="31" width="10" height="8" rx="1.5" fill="#0b1329"/>
                <!-- Колёса -->
                <circle cx="16" cy="48" r="6" fill="#1e293b"/>
                <circle cx="16" cy="48" r="3" fill="#94a3b8"/>
                <circle cx="46" cy="48" r="6" fill="#1e293b"/>
                <circle cx="46" cy="48" r="3" fill="#94a3b8"/>
                <!-- Фары -->
                <rect x="52" y="40" width="3" height="4" rx="1" fill="#FDE68A"/>
              </svg>
            </div>
            <h5>{{ t.adv6Title }}</h5>
            <p>{{ t.adv6Desc }}</p>
          </div>
        </div>
      </div>
    </section>

    <!-- 8. Подвал -->
    <footer id="contacts" class="footer">
      <div class="container footer-grid">
        <div class="foot-brand">
          <h4>QAZAQ PIPE</h4>
          <p>ПОЛИЭТИЛЕНОВЫЕ ТРУБЫ</p>
          <span class="copyright">{{ t.footerRights }}</span>
        </div>
        <div class="foot-contact">
          <span class="icon">📞</span>
          <div>
            <a href="tel:+77027312075">+7 702 731 20 75</a><br>
            <a href="tel:+77780007601">+7 778 000 76 01</a>
          </div>
        </div>
        <div class="foot-contact">
          <span class="icon">✉️</span>
          <div>
            <a href="mailto:samat1979@gmail.com">samat1979@gmail.com</a>
          </div>
        </div>
        <div class="foot-contact">
          <span class="icon">📍</span>
          <div>
            <p>{{ t.footerAddress }}</p>
          </div>
        </div>
        <div>
          <a :href="getWhatsappUrl('Здравствуйте!')" target="_blank" class="btn-foot">{{ t.footerBtn }}</a>
        </div>
      </div>
    </footer>

    <!-- Кнопка «Наверх» — труба с гейзером -->
    <button
      class="scroll-top-btn"
      :class="{ visible: showScrollTop }"
      type="button"
      aria-label="Наверх"
      @click="scrollToTop"
    >
      <span class="geyser-pipe">
        <svg class="pipe-svg" viewBox="0 0 48 64" fill="none" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
          <rect x="16" y="28" width="16" height="32" rx="2" fill="#64748b"/>
          <rect x="18" y="28" width="4" height="32" fill="#94a3b8" opacity="0.45"/>
          <rect x="12" y="24" width="24" height="8" rx="2" fill="#475569"/>
          <rect x="14" y="26" width="20" height="4" rx="1" fill="#334155"/>
          <ellipse cx="24" cy="24" rx="6" ry="3" fill="#0f172a"/>
        </svg>
        <span class="water-jet jet-1"></span>
        <span class="water-jet jet-2"></span>
        <span class="water-jet jet-3"></span>
        <span class="water-jet jet-4"></span>
        <span class="water-drop drop-1"></span>
        <span class="water-drop drop-2"></span>
        <span class="water-drop drop-3"></span>
      </span>
    </button>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'

import rawMaterialsImage from './assets/production/raw_materials.png'
import dryingAndLoadingImage from './assets/production/drying_and_loading.png'
import extrusionImage from './assets/production/extrusion.png'
import formationImage from './assets/production/formation.png'
import coolingImage from './assets/production/cooling.png'
import controlImage from './assets/production/control.png'
import cuttingAndMarkingImage from './assets/production/cutting_and_marking.png'
import finishedGoodsImage from './assets/production/finished_goods.png'

const isMenuOpen = ref(false)
const currentLang = ref('RU') // Текущий язык по умолчанию
const showScrollTop = ref(false)

const onWindowScroll = () => {
  showScrollTop.value = window.scrollY > 400
}

const scrollToTop = () => {
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

// Словарь текстов для переключения RU / KZ
const translations = {
  RU: {
    home: 'ГЛАВНАЯ',
    catalog: 'ПРОДУКЦИЯ',
    calc: 'КАЛЬКУЛЯТОР',
    about: 'О КОМПАНИИ',
    production: 'ПРОИЗВОДСТВО',
    certificates: 'СЕРТИФИКАТЫ',
    advantages: 'ПРЕИМУЩЕСТВА',
    contacts: 'КОНТАКТЫ',
    btnKp: 'ПОЛУЧИТЬ КП',
    heroTitle: 'ПОЛИЭТИЛЕНОВЫЕ ТРУБЫ ОТ ПРОИЗВОДИТЕЛЯ В КАЗАХСТАНЕ',
    heroSubtitle: 'Надёжные решения для водоснабжения, газоснабжения, канализации и промышленных трубопроводов',
    btnHero: 'ПОЛУЧИТЬ КОММЕРЧЕСКОЕ ПРЕДЛОЖЕНИЕ',
    stat1Title: 'С 2010 ГОДА',
    stat1Desc: 'на рынке Казахстана',
    stat2Title: '7 000 ТОНН/ГОД',
    stat2Desc: 'производственная мощность',
    stat3Title: '63–500 ММ',
    stat3Desc: 'диаметры труб',
    stat4Title: 'PE-100 / PE-80',
    stat4Desc: 'основной материал',
    calcTitle: 'КАЛЬКУЛЯТОР РАСЧЕТА ТРУБ',
    calcSubtitle: 'Рассчитайте ориентировочный вес или длину полиэтиленовой трубы',
    calcDiamLabel: 'Диаметр трубы (мм):',
    calcSdrLabel: 'Тип SDR:',
    calcMetersLabel: 'Длина (метров):',
    calcMetersPlaceholder: 'Введите метры',
    calcRes1: 'Примерный вес 1 метра:',
    calcRes2: 'Общий вес заказа:',
    calcBtn: 'ЗАПРОСИТЬ ЦЕНУ ДЛЯ ЭТОГО ОБЪЕМА',
    aboutSub: 'О КОМПАНИИ',
    aboutTitle: 'QAZAQ PIPE —\nОТЕЧЕСТВЕННЫЙ ПРОИЗВОДИТЕЛЬ ПОЛИЭТИЛЕНОВЫХ ТРУБ',
    aboutDesc1: 'ТОО «Qazaq Pipe» работает на рынке Казахстана с 2010 года. Завод расположен в Мангистауской области, г. Актау, на собственной производственной базе с производственным цехом и складскими помещениями для сырья и готовой продукции.',
    aboutDesc2: 'Мы производим полиэтиленовые трубы из ПНД для систем водоснабжения, газоснабжения, канализации и промышленных трубопроводов.',
    aboutBtn: 'ПОДРОБНЕЕ О КОМПАНИИ',
    prodTitle: 'НАША ПРОДУКЦИЯ',
    p1Title: 'ТРУБЫ ДЛЯ ВОДОСНАБЖЕНИЯ',
    p1Desc: 'Ø 63–500 мм',
    p2Title: 'ТРУБЫ ДЛЯ ГАЗОСНАБЖЕНИЯ',
    p2Desc: 'PE-80 / PE-100',
    p3Title: 'ТРУБЫ ДЛЯ КАНАЛИЗАЦИИ',
    p3Desc: 'Напорные и безнапорные',
    p4Title: 'ПРОМЫШЛЕННЫЕ ТРУБОПРОВОДЫ',
    p4Desc: 'Для химически агрессивных сред',
    p5Title: 'ТРУБЫ ДЛЯ ИРРИГАЦИИ',
    p5Desc: 'Мелиорация, орошение, капельный полив',
    p6Title: 'ТРУБЫ-ОБОЛОЧКИ',
    p6Desc: 'Ø 63–500 мм',
    tableTitle: 'ПАРАМЕТРЫ ТРУБ (ТОЛЩИНА СТЕНКИ, ММ)',
    tableTh1: 'Диаметр, мм',
    techTitle: 'ТЕХНОЛОГИЧЕСКАЯ ЛИНИЯ ПРОИЗВОДСТВА',
    techSubtitle: 'Современная экструзионная линия KraussMaffei (Германия)',
    step1Title: 'Сырьё',
    step1Desc: 'Полиэтиленовые гранулы',
    step2Title: 'Сушка и загрузка',
    step2Desc: 'Автоматическая подача гранулята',
    step3Title: 'Экструзия',
    step3Desc: 'Нагрев и плавление полиэтилена',
    step4Title: 'Формирование',
    step4Desc: 'Матричная головка и калибраторы',
    step5Title: 'Охлаждение',
    step5Desc: 'Вакуумная калибровка и водяные ванны',
    step6Title: 'Контроль',
    step6Desc: 'Контроль диаметра и толщины стенки',
    step7Title: 'Резка и маркировка',
    step7Desc: 'Планетарная пила и маркиратор',
    step8Title: 'Готовая продукция',
    step8Desc: 'Складирование и отгрузка',
    certTitle: 'СЕРТИФИКАТЫ И КАЧЕСТВО',
    certSubtitle: 'Официальные документы, подтверждающие соответствие продукции стандартам Республики Казахстан',
    advTitle: 'ПОЧЕМУ ВЫБИРАЮТ QAZAQ PIPE',
    adv1Title: 'СОБСТВЕННОЕ ПРОИЗВОДСТВО',
    adv1Desc: 'Без посредников и лишних звеньев',
    adv2Title: 'НЕМЕЦКОЕ ОБОРУДОВАНИЕ',
    adv2Desc: 'Экструзионная линия KraussMaffei',
    adv3Title: 'СТАБИЛЬНОЕ КАЧЕСТВО',
    adv3Desc: 'Лабораторные испытания и сертификация',
    adv4Title: 'ДО 7 000 ТОНН В ГОД',
    adv4Desc: 'Производственная мощность',
    adv5Title: 'ДИАМЕТРЫ 63–500 ММ',
    adv5Desc: 'Различные SDR и параметры',
    adv6Title: 'ПОСТАВКА ПО КАЗАХСТАНУ',
    adv6Desc: 'Отгрузка продукции с завода в Актау',
    footerRights: '© 2010–2026 ТОО «Qazaq Pipe». Все права защищены.',
    footerAddress: 'Мангистауская область,\nг. Актау, с. Умирзак',
    footerBtn: 'СВЯЗАТЬСЯ С НАМИ'
  },
  KZ: {
    home: 'БАСТЫ БЕТ',
    catalog: 'ӨНІМДЕР',
    calc: 'КАЛЬКУЛЯТОР',
    about: 'КОМПАНИЯ ТУРАЛЫ',
    production: 'ӨНДІРІС',
    certificates: 'СЕРТИФИКАТТАР',
    advantages: 'АРТЫҚШЫЛЫҚТАР',
    contacts: 'БАЙЛАНЫС',
    btnKp: 'КО СҰРАУ',
    heroTitle: 'ҚАЗАҚСТАНДАҒЫ ӨНДІРУШІДЕН ПОЛИЭТИЛЕН ҚҰБЫРЛАРЫ',
    heroSubtitle: 'Сумен жабдықтау, газдандыру, кәріз және өнеркәсіптік құбырларға арналған сенімді шешімдер',
    btnHero: 'КОММЕРЦИЯЛЫҚ ҰСЫНЫС АЛУ',
    stat1Title: '2010 ЖЫЛДАН БЕРІ',
    stat1Desc: 'Қазақстан нарығында',
    stat2Title: '7 000 ТОННА/ЖЫЛ',
    stat2Desc: 'өндірістік қуаттылық',
    stat3Title: '63–500 ММ',
    stat3Desc: 'құбыр диаметрлері',
    stat4Title: 'PE-100 / PE-80',
    stat4Desc: 'негізгі материал',
    calcTitle: 'ҚҰБЫРДЫ ЕСЕПТЕУ КАЛЬКУЛЯТОРЫ',
    calcSubtitle: 'Полиэтилен құбырының шамамен салмағын немесе ұзындығын есептеңіз',
    calcDiamLabel: 'Құбыр диаметрі (мм):',
    calcSdrLabel: 'SDR түрі:',
    calcMetersLabel: 'Ұзындығы (метр):',
    calcMetersPlaceholder: 'Метрді енгізіңіз',
    calcRes1: '1 метрдің шамамен салмағы:',
    calcRes2: 'Тапсырыстың жалпы салмағы:',
    calcBtn: 'БҰЛ КӨЛЕМГЕ БАҒА СҰРАУ',
    aboutSub: 'КОМПАНИЯ ТУРАЛЫ',
    aboutTitle: 'QAZAQ PIPE —\nОТАНДЫҚ ПОЛИЭТИЛЕН ҚҰБЫРЛАРЫН ӨНДІРУШІ',
    aboutDesc1: '«Qazaq Pipe» ЖШС Қазақстан нарығында 2010 жылдан бері жұмыс істейді. Зауыт Маңғыстау облысы, Ақтау қаласында, өндірістік цехы және шикізат пен дайын өнім қоймалары бар өз базасында орналасқан.',
    aboutDesc2: 'Біз сумен жабдықтау, газдандыру, кәріз және өнеркәсіптік құбыр жүйелеріне арналған ПНД полиэтилен құбырларын шығарамыз.',
    aboutBtn: 'КОМПАНИЯ ТУРАЛЫ ТОЛЫҒЫРАҚ',
    prodTitle: 'БІЗДІҢ ӨНІМДЕР',
    p1Title: 'СУМЕН ЖАБДЫҚТАУ ҚҰБЫРЛАРЫ',
    p1Desc: 'Ø 63–500 мм',
    p2Title: 'ГАЗБЕН ЖАБДЫҚТАУ ҚҰБЫРЛАРЫ',
    p2Desc: 'PE-80 / PE-100',
    p3Title: 'КӘРІЗ ҚҰБЫРЛАРЫ',
    p3Desc: 'Қысымды және қысымсыз',
    p4Title: 'ӨНДІРІСТІК ҚҰБЫРЛАР',
    p4Desc: 'Химиялық агрессивті орталар үшін',
    p5Title: 'ИРИГАЦИЯ ҚҰБЫРЛАРЫ',
    p5Desc: 'Мелиорация, суғару, тамшылатып суғару',
    p6Title: 'ҚАБЫҚ ҚҰБЫРЛАР',
    p6Desc: 'Ø 63–500 мм',
    tableTitle: 'ҚҰБЫР ПАРАМЕТРЛЕРІ (ҚАБЫРҒА ҚАЛЫҢДЫҒЫ, ММ)',
    tableTh1: 'Диаметрі, мм',
    techTitle: 'ӨНДІРІСТІК ТЕХНОЛОГИЯЛЫҚ ЖЕЛІ',
    techSubtitle: 'Заманауи KraussMaffei экструзиялық желісі (Германия)',
    step1Title: 'Шикізат',
    step1Desc: 'Полиэтилен түйіршіктері',
    step2Title: 'Кептіру және тиеу',
    step2Desc: 'Түйіршіктерді автоматты түрде беру',
    step3Title: 'Экструзия',
    step3Desc: 'Полиэтиленді қыздыру және балқыту',
    step4Title: 'Қалыптастыру',
    step4Desc: 'Матрицалық бас және калибраторлар',
    step5Title: 'Салқындату',
    step5Desc: 'Вакуумдық калибрлеу және су ванналары',
    step6Title: 'Бақылау',
    step6Desc: 'Диаметр мен қабырға қалыңдығын бақылау',
    step7Title: 'Кесу және таңбалау',
    step7Desc: 'Планетарлық ара және таңбалау құрылғысы',
    step8Title: 'Дайын өнім',
    step8Desc: 'Қоймаға жинау және жөнелту',
    certTitle: 'СЕРТИФИКАТТАР ЖӘНЕ САПА',
    certSubtitle: 'Өнімнің Қазақстан Республикасы стандарттарына сәйкестігін растайтын ресми құжаттар',
    advTitle: 'НЕГЕ QAZAQ PIPE ТАҢДАЙДЫ',
    adv1Title: 'БӨЛЕК ӨНДІРІС',
    adv1Desc: 'Дельфиктерсіз және артық буындарсыз',
    adv2Title: 'НЕМІС ЖАБДЫҚТАРЫ',
    adv2Desc: 'KraussMaffei экструзиялық желісі',
    adv3Title: 'ТҰРАҚТЫ САПА',
    adv3Desc: 'Зертханалық сынақтар және сертификаттау',
    adv4Title: 'ЖЫЛЫНА 7 000 ТОННАҒА ДЕЙІН',
    adv4Desc: 'Өндірістік қуаттылық',
    adv5Title: '63–500 ММ ДИАМЕТРЛЕР',
    adv5Desc: 'Әртүрлі SDR және параметрлер',
    adv6Title: 'ҚАЗАҚСТАН БОЙЫНША ЖЕТКІЗУ',
    adv6Desc: 'Ақтаудағы зауыттан өнімді жөнелту',
    footerRights: '© 2010–2026 «Qazaq Pipe» ЖШС. Барлық құқықтар қорғалған.',
    footerAddress: 'Маңғыстау облысы,\nАқтау қ., Өмірзақ қ.а.',
    footerBtn: 'БІЗБЕН БАЙЛАНЫСУ'
  }
}

// Актуальный объект перевода
const t = computed(() => translations[currentLang.value])

// Функция переключения языка
const toggleLanguage = () => {
  currentLang.value = currentLang.value === 'RU' ? 'KZ' : 'RU'
}

// Сертификаты (PDF) — файлы в public/certificates/
// lang — язык документа (бейдж на карточке)
const certificatesList = [
  {
    file: '/certificates/cert-gas-ru.pdf',
    lang: 'RU',
    titleRu: 'Сертификат соответствия — газ',
    titleKz: 'Сәйкестік сертификаты — газ',
    descRu: 'Трубы ПЭ для газопроводов · № KZ.7500533.01.01.00740 · до 10.04.2027',
    descKz: 'Газ құбырларына арналған ПЭ құбырлар · № KZ.7500533.01.01.00740 · 10.04.2027 дейін'
  },
  {
    file: '/certificates/cert-gas-kz.pdf',
    lang: 'KZ',
    titleRu: 'Сертификат соответствия — газ',
    titleKz: 'Сәйкестік сертификаты — газ',
    descRu: 'Трубы ПЭ для газопроводов · № KZ.7500533.01.01.00740 · до 10.04.2027',
    descKz: 'Газ құбырларына арналған ПЭ құбырлар · № KZ.7500533.01.01.00740 · 10.04.2027 дейін'
  },
  {
    file: '/certificates/cert-water-ru.pdf',
    lang: 'RU',
    titleRu: 'Сертификат соответствия — вода',
    titleKz: 'Сәйкестік сертификаты — су',
    descRu: 'Трубы ПЭ-100 для питьевого водоснабжения · ГОСТ 18599-2001 · № KZ.7500533.01.01.01418',
    descKz: 'Ауыз суға арналған ПЭ-100 құбырлар · ГОСТ 18599-2001 · № KZ.7500533.01.01.01418'
  },
  {
    file: '/certificates/cert-water-kz.pdf',
    lang: 'KZ',
    titleRu: 'Сертификат соответствия — вода',
    titleKz: 'Сәйкестік сертификаты — су',
    descRu: 'Трубы ПЭ-100 для питьевого водоснабжения · ГОСТ 18599-2001 · № KZ.7500533.01.01.01418',
    descKz: 'Ауыз суға арналған ПЭ-100 құбырлар · ГОСТ 18599-2001 · № KZ.7500533.01.01.01418'
  },
  {
    file: '/certificates/cert-sanitary.pdf',
    lang: 'RU',
    titleRu: 'Свидетельство о гос. регистрации',
    titleKz: 'Мемлекеттік тіркеу куәлігі',
    descRu: 'ЕАЭС · ПЭ-100 для питьевого водоснабжения · № KG.11.01.09.013.E.006528.12.24',
    descKz: 'ЕАЭО · Ауыз суға арналған ПЭ-100 · № KG.11.01.09.013.E.006528.12.24'
  }
]

const certificates = computed(() =>
  certificatesList.map((c) => ({
    file: c.file,
    lang: c.lang,
    title: currentLang.value === 'RU' ? c.titleRu : c.titleKz,
    desc: currentLang.value === 'RU' ? c.descRu : c.descKz
  }))
)

// Карусель технологического процесса
const productionSteps = computed(() => [
  {
    number: '01',
    title: t.value.step1Title,
    description: t.value.step1Desc,
    image: rawMaterialsImage
  },
  {
    number: '02',
    title: t.value.step2Title,
    description: t.value.step2Desc,
    image: dryingAndLoadingImage
  },
  {
    number: '03',
    title: t.value.step3Title,
    description: t.value.step3Desc,
    image: extrusionImage
  },
  {
    number: '04',
    title: t.value.step4Title,
    description: t.value.step4Desc,
    image: formationImage
  },
  {
    number: '05',
    title: t.value.step5Title,
    description: t.value.step5Desc,
    image: coolingImage
  },
  {
    number: '06',
    title: t.value.step6Title,
    description: t.value.step6Desc,
    image: controlImage
  },
  {
    number: '07',
    title: t.value.step7Title,
    description: t.value.step7Desc,
    image: cuttingAndMarkingImage
  },
  {
    number: '08',
    title: t.value.step8Title,
    description: t.value.step8Desc,
    image: finishedGoodsImage
  }
])

const visibleSlides = ref(4)
const carouselWindow = ref(null)
const carouselWidth = ref(0)
const carouselGap = 20
const currentSlide = ref(0)
const isResettingCarousel = ref(false)
let carouselTimer = null
let resetTimer = null

const updateVisibleSlides = () => {
  carouselWidth.value = carouselWindow.value?.clientWidth || 0
  if (window.innerWidth <= 576) {
    visibleSlides.value = 1
  } else if (window.innerWidth <= 992) {
    visibleSlides.value = 2
  } else {
    visibleSlides.value = 4
  }

  currentSlide.value = Math.min(
    currentSlide.value,
    productionSteps.value.length - 1
  )
}

const carouselTrackStyle = computed(() => {
  const width = carouselWidth.value || carouselWindow.value?.clientWidth || 0
  const slideWidth = visibleSlides.value
    ? (width - carouselGap * (visibleSlides.value - 1)) / visibleSlides.value
    : width
  const step = slideWidth + carouselGap

  return {
    transform: `translateX(-${currentSlide.value * step}px)`,
    gap: `${carouselGap}px`
  }
})

const carouselSlides = computed(() => {
  const count = visibleSlides.value
  return [
    ...productionSteps.value,
    ...productionSteps.value.slice(0, count)
  ]
})

const activeDot = computed(() => {
  return currentSlide.value % productionSteps.value.length
})

const nextSlide = () => {
  currentSlide.value += 1

  if (currentSlide.value === productionSteps.value.length) {
    clearTimeout(resetTimer)

    resetTimer = setTimeout(() => {
      isResettingCarousel.value = true
      currentSlide.value = 0

      requestAnimationFrame(() => {
        requestAnimationFrame(() => {
          isResettingCarousel.value = false
        })
      })
    }, 650)
  }
}

const prevSlide = () => {
  if (currentSlide.value === 0) {
    isResettingCarousel.value = true
    currentSlide.value = productionSteps.value.length

    requestAnimationFrame(() => {
      requestAnimationFrame(() => {
        isResettingCarousel.value = false
        currentSlide.value = productionSteps.value.length - 1
      })
    })
    return
  }

  currentSlide.value -= 1
}

const goToSlide = (index) => {
  clearTimeout(resetTimer)
  isResettingCarousel.value = false
  currentSlide.value = index
  restartCarousel()
}

const restartCarousel = () => {
  clearInterval(carouselTimer)
  carouselTimer = setInterval(nextSlide, 3500)
}

onMounted(() => {
  updateVisibleSlides()
  window.addEventListener('resize', updateVisibleSlides)
  window.addEventListener('scroll', onWindowScroll, { passive: true })
  onWindowScroll()
  restartCarousel()
})

onBeforeUnmount(() => {
  window.removeEventListener('resize', updateVisibleSlides)
  window.removeEventListener('scroll', onWindowScroll)
  clearInterval(carouselTimer)
  clearTimeout(resetTimer)
})

// Калькулятор
const calcDiameter = ref(110)
const calcSdr = ref(11)
const calcMeters = ref(100)

const diameterOptions = [63, 90, 110, 160, 225, 315, 400, 500]
const sdrOptions = [11, 17, 21, 26]

const weightTable = {
  63:  { 11: 1.04, 17: 0.70, 21: 0.58, 26: 0.47 },
  90:  { 11: 2.11, 17: 1.43, 21: 1.18, 26: 0.98 },
  110: { 11: 3.17, 17: 2.14, 21: 1.76, 26: 1.43 },
  160: { 11: 6.71, 17: 4.51, 21: 3.71, 26: 3.03 },
  225: { 11: 13.3, 17: 8.93, 21: 7.32, 26: 5.97 },
  315: { 11: 26.0, 17: 17.5, 21: 14.3, 26: 11.7 },
  400: { 11: 41.8, 17: 28.1, 21: 23.0, 26: 18.8 },
  500: { 11: 65.2, 17: 43.9, 21: 35.9, 26: 29.3 }
}

const weightPerMeter = computed(() => {
  return weightTable[calcDiameter.value]?.[calcSdr.value] || 0
})

const totalWeight = computed(() => {
  return Math.round(weightPerMeter.value * (calcMeters.value || 0))
})

const getWhatsappUrl = (msg) => {
  return `https://wa.me/77027312075?text=${encodeURIComponent(msg)}`
}

const closeMenu = () => {
  isMenuOpen.value = false
}
</script>

<style>
/* Google Fonts — Manrope (заголовки) + Inter (текст) */
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&family=Manrope:wght@600;700;800&display=swap');

/* 1. Глобальный сброс и типографика */
*, *::before, *::after {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

html, body {
  margin: 0 !important;
  padding: 0 !important;
  width: 100% !important;
  max-width: 100% !important;
  overflow-x: hidden;
  background-color: #ffffff;
  color: #1e293b;
  scroll-behavior: smooth;
  font-family: 'Inter', system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  font-size: 15px;
  line-height: 1.55;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-rendering: optimizeLegibility;
  font-feature-settings: "kern" 1, "liga" 1;
}

#app, .app {
  width: 100% !important;
  max-width: 100% !important;
  margin: 0 !important;
  padding: 0 !important;
}

.container { 
  max-width: 1240px !important; 
  margin: 0 auto !important; 
  padding: 0 15px !important; 
  width: 100% !important;
}

.navbar, .hero, .section-calculator, .section-about, .section-gray, .section-white, .section-certificates, .section-dark, .footer {
  width: 100% !important;
}

/* 2. Navbar & Навигация */
.navbar { 
  background: #0b1329; 
  color: white; 
  position: sticky; 
  top: 0; 
  z-index: 1000; 
  padding: 12px 0; 
  border-bottom: 1px solid rgba(255,255,255,0.1); 
}
.navbar-content { 
  display: flex; 
  justify-content: space-between; 
  align-items: center; 
}

.logo {
  display: flex;
  flex-direction: column;
  justify-content: center;
}
.logo h1 {
  font-family: 'Manrope', sans-serif;
  color: #ffffff;
  font-size: 21px;
  font-weight: 800;
  letter-spacing: 1px;
  line-height: 1.1;
  margin: 0;
}
.logo span {
  font-family: 'Inter', sans-serif;
  font-size: 10px;
  color: #94a3b8;
  display: block;
  font-weight: 600;
  letter-spacing: 1.2px;
  text-transform: uppercase;
  margin: 0;
}

.nav-links { display: flex; gap: 18px; align-items: center; }
.nav-links a {
  text-decoration: none;
  color: #cbd5e1;
  font-family: 'Inter', sans-serif;
  font-size: 12px;
  font-weight: 600;
  letter-spacing: 0.4px;
  transition: color 0.2s;
}
.nav-links a:hover { color: #38bdf8; }

.nav-right { display: flex; align-items: center; gap: 12px; }
.btn-kp {
  font-family: 'Manrope', sans-serif;
  background: #0284c7;
  color: white;
  padding: 8px 16px;
  border-radius: 6px;
  text-decoration: none;
  font-weight: 700;
  font-size: 12px;
  letter-spacing: 0.3px;
  transition: background 0.2s, transform 0.15s;
}
.btn-kp:hover { background: #0369a1; transform: translateY(-1px); }

/* Кнопка переключения языка */
.lang-switch {
  font-family: 'Manrope', sans-serif;
  font-size: 12px;
  font-weight: 700;
  color: #ffffff;
  background: #1e293b;
  cursor: pointer;
  border: 1px solid #334155;
  padding: 6px 12px;
  border-radius: 6px;
  transition: all 0.2s ease;
  letter-spacing: 0.5px;
}
.lang-switch:hover {
  background: #0284c7;
  border-color: #0284c7;
}

.mobile-kp { display: none; }

/* Кнопка бургер-меню */
.burger-btn {
  display: none;
  flex-direction: column;
  justify-content: space-between;
  width: 26px;
  height: 20px;
  background: transparent;
  border: none;
  cursor: pointer;
  padding: 0;
  z-index: 1001;
}
.burger-btn span {
  width: 100%;
  height: 2px;
  background-color: #ffffff;
  border-radius: 2px;
  transition: all 0.3s ease;
}

.burger-btn span.open:nth-child(1) {
  transform: translateY(9px) rotate(45deg);
}
.burger-btn span.open:nth-child(2) {
  opacity: 0;
}
.burger-btn span.open:nth-child(3) {
  transform: translateY(-9px) rotate(-45deg);
}

/* 3. Hero Баннер */
.hero { 
  background: linear-gradient(rgba(11,19,41,0.85), rgba(11,19,41,0.85)), 
              url('https://images.unsplash.com/photo-1542013936693-884638332954?auto=format&fit=crop&w=1600&q=80') center/cover no-repeat; 
  color: white; 
  padding: 60px 0 0; 
  position: relative; 
  text-align: center; 
}
.hero-container { 
  display: flex; 
  flex-direction: column; 
  align-items: center; 
}
.hero-text-block {
  margin-bottom: 40px;
  max-width: 820px;
  width: 100%;
  text-align: center;
  display: flex;
  flex-direction: column;
  align-items: center;
}
.hero-text-block h2 {
  font-family: 'Manrope', sans-serif;
  font-size: 36px;
  font-weight: 800;
  line-height: 1.22;
  margin-bottom: 16px;
  letter-spacing: -0.5px;
  text-align: center;
  width: 100%;
}
.hero-text-block .subtitle {
  font-family: 'Inter', sans-serif;
  font-size: 16px;
  font-weight: 400;
  color: #cbd5e1;
  margin: 0 auto 28px;
  line-height: 1.65;
  max-width: 640px;
  text-align: center;
  width: 100%;
}
.btn-hero {
  font-family: 'Manrope', sans-serif;
  background: #0284c7;
  color: white;
  padding: 14px 28px;
  border-radius: 8px;
  text-decoration: none;
  font-weight: 700;
  font-size: 14px;
  display: inline-block;
  letter-spacing: 0.3px;
  text-align: center;
  margin: 0 auto;
  transition: background 0.2s, transform 0.15s, box-shadow 0.2s;
}
.btn-hero:hover {
  background: #0369a1;
  transform: translateY(-2px);
  box-shadow: 0 8px 20px rgba(2, 132, 199, 0.35);
}

.hero-stats-bar { 
  display: grid; 
  grid-template-columns: repeat(4, 1fr); 
  background: rgba(15, 23, 42, 0.9); 
  border: 1px solid rgba(255,255,255,0.1); 
  border-radius: 8px 8px 0 0; 
  margin-top: 40px; 
  padding: 20px; 
  width: 100%; 
  text-align: left; 
}
.stat-box { display: flex; align-items: center; gap: 14px; border-right: 1px solid rgba(255,255,255,0.1); padding: 0 10px; }
.stat-box:last-child { border-right: none; }

.stat-icon { 
  display: flex; 
  align-items: center; 
  justify-content: center; 
  color: #0284c7; 
  flex-shrink: 0;
}

.stat-box strong {
  font-family: 'Manrope', sans-serif;
  display: block;
  font-size: 15px;
  font-weight: 700;
  color: white;
  letter-spacing: 0.2px;
}
.stat-box span { font-size: 11px; color: #94a3b8; }

/* 4. Калькулятор */
.section-calculator { padding: 80px 0; background: #f1f5f9; }
.calc-card {
  background: white;
  border-radius: 14px;
  border: 1px solid #e2e8f0;
  padding: 42px 48px;
  max-width: 980px;
  margin: 0 auto;
  box-shadow: 0 8px 24px -4px rgba(15, 23, 42, 0.08);
}
.calc-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 24px;
  margin-bottom: 28px;
}
.calc-field { display: flex; flex-direction: column; gap: 10px; }
.calc-field label {
  font-family: 'Inter', sans-serif;
  font-size: 13px;
  font-weight: 600;
  color: #334155;
  letter-spacing: 0.2px;
}
.calc-field select,
.calc-field input {
  padding: 14px 16px;
  border: 1.5px solid #cbd5e1;
  border-radius: 8px;
  font-size: 16px;
  font-family: 'Inter', sans-serif;
  outline: none;
  color: #0b1329;
  background: #fff;
  min-height: 52px;
  transition: border-color 0.2s, box-shadow 0.2s;
}
.calc-field select:focus,
.calc-field input:focus {
  border-color: #0284c7;
  box-shadow: 0 0 0 3px rgba(2, 132, 199, 0.15);
}

.calc-result {
  background: #f8fafc;
  padding: 28px 24px;
  border-radius: 12px;
  border: 1px solid #e2e8f0;
  display: flex;
  flex-direction: column;
  gap: 14px;
  align-items: center;
  text-align: center;
}
.res-item {
  font-family: 'Inter', sans-serif;
  font-size: 16px;
  color: #475569;
}
.res-item strong {
  font-family: 'Manrope', sans-serif;
  color: #0b1329;
  margin-left: 8px;
  font-size: 18px;
}
.res-item.main-res {
  font-size: 18px;
  font-weight: 600;
}
.res-item.main-res strong {
  color: #0284c7;
  font-size: 26px;
  font-weight: 800;
}
.btn-calc {
  font-family: 'Manrope', sans-serif;
  background: #0284c7;
  color: white;
  padding: 16px 32px;
  border-radius: 10px;
  text-decoration: none;
  font-weight: 700;
  font-size: 15px;
  margin-top: 8px;
  letter-spacing: 0.3px;
  transition: background 0.2s, transform 0.15s, box-shadow 0.2s;
  display: inline-block;
}
.btn-calc:hover {
  background: #0369a1;
  transform: translateY(-2px);
  box-shadow: 0 8px 20px rgba(2, 132, 199, 0.3);
}

/* 5. О Компании */
.section-about { padding: 70px 0; background: #ffffff; }
.grid-about { display: grid; grid-template-columns: 1fr 1fr; gap: 50px; align-items: center; }
.sub-label { font-size: 12px; font-weight: 800; color: #64748b; letter-spacing: 1.2px; }
.about-text h3 {
  font-family: 'Manrope', sans-serif;
  font-size: 26px;
  font-weight: 800;
  color: #0b1329;
  margin: 10px 0 20px;
  line-height: 1.3;
  white-space: pre-line;
  letter-spacing: -0.4px;
}
.about-text p { font-size: 14px; color: #475569; line-height: 1.65; margin-bottom: 15px; }
.btn-about { display: inline-block; background: #0b1329; color: white; padding: 10px 20px; text-decoration: none; font-size: 12px; font-weight: 700; border-radius: 4px; margin-top: 10px; letter-spacing: 0.5px; }

.about-img-box { 
  background: #f8fafc; 
  height: 340px; 
  border-radius: 8px; 
  overflow: hidden; 
  display: flex; 
  align-items: center; 
  justify-content: center; 
  border: 1px solid #e2e8f0; 
  box-shadow: 0 4px 6px -1px rgba(0,0,0,0.05);
}
.about-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.4s ease;
  image-rendering: -webkit-optimize-contrast;
}
.about-img-box:hover .about-img {
  transform: scale(1.03);
}

/* 6. Продукция и Таблица */
.section-gray { background: #f8fafc; padding: 70px 0; }
.section-title {
  font-family: 'Manrope', sans-serif;
  font-size: 24px;
  font-weight: 800;
  color: #0b1329;
  letter-spacing: -0.3px;
  margin-bottom: 32px;
}
.section-title.center { text-align: center; }
.section-subtitle {
  font-family: 'Inter', sans-serif;
  color: #64748b;
  font-size: 14px;
  font-weight: 400;
  margin-top: -18px;
  margin-bottom: 40px;
  line-height: 1.5;
}
.section-subtitle.center { text-align: center; }

.product-cards-grid { display: grid; grid-template-columns: repeat(6, 1fr); gap: 15px; margin-bottom: 50px; }
.prod-card { background: white; border-radius: 6px; padding: 15px; text-align: center; border: 1px solid #e2e8f0; }

.animated-card {
  transition: transform 0.35s cubic-bezier(0.16, 1, 0.3, 1), 
              box-shadow 0.35s cubic-bezier(0.16, 1, 0.3, 1);
  will-change: transform, box-shadow;
}
.animated-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 20px 30px -10px rgba(11, 19, 41, 0.12);
}
.animated-card:hover .prod-img {
  transform: scale(1.08);
}

.prod-img-container { 
  height: 90px; 
  border-radius: 4px; 
  background: #0b1329; 
  overflow: hidden; 
  margin-bottom: 12px; 
  display: flex;
  align-items: center;
  justify-content: center;
}
.prod-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.4s ease;
}

.prod-card h4 {
  font-family: 'Manrope', sans-serif;
  font-size: 12px;
  font-weight: 800;
  color: #0b1329;
  margin-bottom: 4px;
  letter-spacing: 0.2px;
}
.prod-card p { font-size: 10px; color: #64748b; }

.table-container { background: white; padding: 25px; border-radius: 8px; border: 1px solid #e2e8f0; }
.table-container h4 { font-size: 14px; font-weight: 800; margin-bottom: 15px; color: #0b1329; }
.table-wrapper { overflow-x: auto; }
.spec-table { width: 100%; border-collapse: collapse; text-align: center; font-size: 13px; min-width: 600px; }
.spec-table th, .spec-table td { padding: 10px; border: 1px solid #e2e8f0; }
.spec-table th { background: #0b1329; color: white; font-weight: 700; }

/* 7. Производство — анимированная карусель */
.section-white { padding: 70px 0; background: white; }

.production-carousel {
  position: relative;
  display: flex;
  align-items: center;
  gap: 12px;
  width: 100%;
}

.carousel-window {
  width: 100%;
  overflow: hidden;
  padding: 12px 2px 20px;
}

.carousel-track {
  display: flex;
  transition: transform 0.6s cubic-bezier(0.22, 1, 0.36, 1);
  will-change: transform;
}

.carousel-track.no-transition {
  transition: none;
}

.production-slide {
  box-sizing: border-box;
  flex: 0 0 calc(100% / 4);
  width: calc(100% / 4);
  min-width: 0;
  margin: 0;
  flex-basis: calc((100% - 60px) / 4);
}

.step-item {
  background: #f8fafc;
  border-radius: 8px;
  padding: 25px 20px;
  position: relative;
  border: 1px solid #e2e8f0;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  min-height: 210px;
}

.production-slide:hover {
  transform: translateY(-8px);
  box-shadow: 0 20px 30px -10px rgba(11, 19, 41, 0.12);
}

.production-image-box {
  width: 100%;
  height: 120px;
  margin-bottom: 16px;
  border-radius: 7px;
  overflow: hidden;
  background: #e0f2fe;
}

.production-image {
  width: 100%;
  height: 100%;
  display: block;
  object-fit: cover;
  transition: transform 0.5s ease;
}

.production-slide:hover .production-image {
  transform: scale(1.05);
}

.step-num {
  position: absolute;
  top: 20px;
  right: 20px;
  z-index: 2;
  font-size: 16px;
  font-weight: 900;
  color: #ffffff;
  text-shadow: 0 1px 4px rgba(0, 0, 0, 0.5);
}

.step-item h5 {
  font-size: 13px;
  font-weight: 800;
  color: #0b1329;
  margin-bottom: 6px;
  letter-spacing: 0.2px;
}

.step-item p {
  font-size: 11px;
  color: #64748b;
  line-height: 1.4;
}

.carousel-arrow {
  flex: 0 0 42px;
  width: 42px;
  height: 42px;
  border: 1px solid #dbe4ee;
  border-radius: 50%;
  background: white;
  color: #0284c7;
  font-size: 30px;
  line-height: 1;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease;
  z-index: 3;
}

.carousel-arrow:hover {
  background: #0284c7;
  color: white;
  border-color: #0284c7;
  transform: scale(1.05);
}

.carousel-dots {
  display: flex;
  justify-content: center;
  gap: 7px;
  margin-top: 8px;
}

.carousel-dot {
  width: 7px;
  height: 7px;
  padding: 0;
  border: none;
  border-radius: 50%;
  background: #cbd5e1;
  cursor: pointer;
  transition: all 0.25s ease;
}

.carousel-dot.active {
  width: 22px;
  border-radius: 10px;
  background: #0284c7;
}

/* 8. Преимущества */
/* Сертификаты */
.section-certificates {
  padding: 70px 0;
  background: #f8fafc;
}
.cert-grid {
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  gap: 16px;
}
.cert-card {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 12px;
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  padding: 20px 16px;
  text-decoration: none;
  color: inherit;
  transition: transform 0.3s ease, box-shadow 0.3s ease, border-color 0.2s;
  min-height: 200px;
}
.cert-card:hover {
  transform: translateY(-6px);
  box-shadow: 0 12px 28px rgba(15, 23, 42, 0.1);
  border-color: #bae6fd;
}
.cert-icon {
  width: 48px;
  height: 48px;
  display: flex;
  align-items: center;
  justify-content: center;
}
.cert-body {
  flex: 1;
  width: 100%;
}
.cert-lang {
  display: inline-block;
  font-family: 'Manrope', sans-serif;
  font-size: 10px;
  font-weight: 800;
  letter-spacing: 0.8px;
  padding: 3px 8px;
  border-radius: 4px;
  margin-bottom: 8px;
}
.cert-lang.lang-ru {
  background: #e0f2fe;
  color: #0369a1;
}
.cert-lang.lang-kz {
  background: #dcfce7;
  color: #15803d;
}
.cert-card h5 {
  font-family: 'Manrope', sans-serif;
  font-size: 13px;
  font-weight: 800;
  color: #0b1329;
  line-height: 1.35;
  margin-bottom: 6px;
}
.cert-card p {
  font-family: 'Inter', sans-serif;
  font-size: 11px;
  color: #64748b;
  line-height: 1.4;
}
.cert-download {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-family: 'Manrope', sans-serif;
  font-size: 11px;
  font-weight: 700;
  color: #0284c7;
  letter-spacing: 0.3px;
  margin-top: auto;
}
.cert-card:hover .cert-download {
  color: #0369a1;
}

.section-dark { background: #0b1329; color: white; padding: 70px 0; }
.white-text { color: white !important; }
.adv-grid { display: grid; grid-template-columns: repeat(6, 1fr); gap: 20px; text-align: center; }
.adv-icon {
  width: 64px;
  height: 64px;
  margin: 0 auto 14px;
  display: flex;
  align-items: center;
  justify-content: center;
}
.adv-icon svg {
  width: 48px;
  height: 48px;
  display: block;
  filter: drop-shadow(0 4px 8px rgba(0,0,0,0.25));
}
.adv-item h5 {
  font-family: 'Manrope', sans-serif;
  font-size: 12px;
  font-weight: 800;
  margin-bottom: 6px;
  letter-spacing: 0.4px;
  text-transform: uppercase;
  color: #ffffff;
}
.adv-item p {
  font-family: 'Inter', sans-serif;
  font-size: 11px;
  color: #94a3b8;
  line-height: 1.4;
}

/* 9. Подвал */
.footer {
  background: #070d1c;
  color: #94a3b8;
  padding: 48px 0;
  border-top: 1px solid rgba(255,255,255,0.08);
}
.footer-grid {
  display: grid;
  grid-template-columns: 1.5fr 1fr 1fr 1fr 1fr;
  gap: 24px;
  align-items: center;
  font-size: 13px;
  font-family: 'Inter', sans-serif;
}
.foot-brand h4 {
  font-family: 'Manrope', sans-serif;
  color: white;
  font-size: 17px;
  font-weight: 800;
  letter-spacing: 0.6px;
}
.copyright {
  font-size: 11px;
  color: #475569;
  display: block;
  margin-top: 10px;
}
.foot-contact { display: flex; gap: 10px; align-items: center; }
.foot-contact a {
  color: white;
  text-decoration: none;
  font-weight: 600;
  font-family: 'Inter', sans-serif;
}
.btn-foot {
  font-family: 'Manrope', sans-serif;
  background: #0284c7;
  color: white;
  padding: 11px 18px;
  border-radius: 6px;
  text-decoration: none;
  font-weight: 700;
  font-size: 12px;
  display: inline-block;
  letter-spacing: 0.3px;
  transition: background 0.2s, transform 0.15s;
}
.btn-foot:hover {
  background: #0369a1;
  transform: translateY(-1px);
}

/* ==========================================================
   АДАПТИВНОСТЬ ДЛЯ МОБИЛЬНЫХ УСТРОЙСТВ
   ========================================================== */
@media (max-width: 992px) {
  .production-slide {
    flex-basis: calc((100% - 20px) / 2);
  }

  .logo h1 { font-size: 16px; }
  .logo span { font-size: 9px; }

  .burger-btn { display: flex; }
  .desktop-kp { display: none; }
  .mobile-kp { display: inline-block; text-align: center; width: 100%; margin-top: 15px; }

  .nav-links {
    position: fixed;
    top: 55px;
    left: 0;
    width: 100%;
    background: #0b1329;
    flex-direction: column;
    padding: 20px;
    gap: 15px;
    border-bottom: 1px solid rgba(255,255,255,0.1);
    display: none;
    box-shadow: 0 10px 15px -3px rgba(0,0,0,0.3);
  }
  .nav-links.nav-active { display: flex; }

  .hero-stats-bar { grid-template-columns: repeat(2, 1fr); gap: 15px; }
  .stat-box { border-right: none; }
  .grid-about { grid-template-columns: 1fr; gap: 30px; }
  .product-cards-grid { grid-template-columns: repeat(3, 1fr); }
  .cert-grid { grid-template-columns: repeat(3, 1fr); }
  .adv-grid { grid-template-columns: repeat(3, 1fr); }
  .footer-grid { grid-template-columns: 1fr; gap: 20px; }
  .calc-grid { grid-template-columns: 1fr; }
}

@media (max-width: 576px) {
  .production-carousel {
    gap: 6px;
  }

  .production-slide {
    flex-basis: 100%;
  }

  .carousel-arrow {
    flex: 0 0 36px;
    width: 36px;
    height: 36px;
    font-size: 25px;
  }

  .hero-text-block h2 { font-size: 24px; }
  .hero-stats-bar { grid-template-columns: 1fr; }
  .product-cards-grid { grid-template-columns: repeat(2, 1fr); }
  .cert-grid { grid-template-columns: 1fr; }
  .adv-grid { grid-template-columns: repeat(2, 1fr); }
}

/* ===== Кнопка «Наверх» — труба с гейзером ===== */
.scroll-top-btn {
  position: fixed;
  right: 22px;
  bottom: 28px;
  z-index: 1100;
  width: 64px;
  height: 72px;
  border: none;
  background: transparent;
  cursor: pointer;
  padding: 0;
  opacity: 0;
  visibility: hidden;
  transform: translateY(16px) scale(0.9);
  transition: opacity 0.3s ease, visibility 0.3s ease, transform 0.3s ease;
  -webkit-tap-highlight-color: transparent;
}
.scroll-top-btn.visible {
  opacity: 1;
  visibility: visible;
  transform: translateY(0) scale(1);
}
.scroll-top-btn:hover {
  transform: translateY(-4px) scale(1.05);
}
.scroll-top-btn:active {
  transform: translateY(0) scale(0.98);
}

.geyser-pipe {
  position: relative;
  display: block;
  width: 48px;
  height: 64px;
  margin: 0 auto;
  filter: drop-shadow(0 6px 12px rgba(15, 23, 42, 0.35));
}
.pipe-svg {
  position: relative;
  z-index: 2;
  width: 48px;
  height: 64px;
  display: block;
}

.water-jet {
  position: absolute;
  left: 50%;
  bottom: 40px;
  width: 5px;
  border-radius: 999px;
  background: linear-gradient(to top, #38bdf8, #7dd3fc, transparent);
  transform: translateX(-50%);
  opacity: 0;
  pointer-events: none;
  z-index: 1;
}
.scroll-top-btn.visible .water-jet {
  animation: geyserJet 1.4s ease-out infinite;
}
.jet-1 { height: 22px; margin-left: 0; animation-delay: 0s; }
.jet-2 { height: 28px; margin-left: -6px; width: 4px; animation-delay: 0.15s; }
.jet-3 { height: 26px; margin-left: 6px; width: 4px; animation-delay: 0.25s; }
.jet-4 { height: 18px; margin-left: 0; width: 3px; animation-delay: 0.4s; }

@keyframes geyserJet {
  0% {
    opacity: 0;
    transform: translateX(-50%) translateY(8px) scaleY(0.3);
  }
  25% { opacity: 1; }
  70% { opacity: 0.85; }
  100% {
    opacity: 0;
    transform: translateX(-50%) translateY(-28px) scaleY(1.2);
  }
}

.water-drop {
  position: absolute;
  left: 50%;
  bottom: 48px;
  width: 5px;
  height: 5px;
  border-radius: 50%;
  background: #7dd3fc;
  opacity: 0;
  pointer-events: none;
  z-index: 3;
}
.scroll-top-btn.visible .water-drop {
  animation: geyserDrop 1.6s ease-in infinite;
}
.drop-1 { margin-left: -10px; animation-delay: 0.2s; }
.drop-2 { margin-left: 8px; animation-delay: 0.45s; width: 4px; height: 4px; }
.drop-3 { margin-left: -2px; animation-delay: 0.7s; width: 3px; height: 3px; }

@keyframes geyserDrop {
  0% {
    opacity: 0;
    transform: translate(-50%, 0) scale(0.5);
  }
  20% { opacity: 1; }
  100% {
    opacity: 0;
    transform: translate(-50%, -36px) scale(1);
  }
}

@media (max-width: 576px) {
  .scroll-top-btn {
    right: 14px;
    bottom: 18px;
    width: 56px;
    height: 64px;
  }
  .geyser-pipe,
  .pipe-svg {
    width: 42px;
    height: 56px;
  }
}

</style>